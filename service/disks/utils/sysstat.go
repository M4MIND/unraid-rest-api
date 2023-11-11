package utils

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unraid-rest-api/service/disks/types"
)

type DisksUtils struct {
}

func NewDiskUtils() DisksUtils {
	return DisksUtils{}
}

func parse(line string) (diskRawStats types.SysstatRaw) {
	diskRawStats = types.SysstatRaw{}
	fields := strings.Fields(line)

	for i := 0; i < len(fields); i++ {
		field := fields[i]
		switch i {
		case 0:
			major, _ := strconv.ParseInt(field, 10, strconv.IntSize)
			diskRawStats.Major = int(major)
		case 1:
			minor, _ := strconv.ParseInt(field, 10, strconv.IntSize)
			diskRawStats.Minor = int(minor)
		case 2:
			diskRawStats.Name = fields[2]
		case 3:
			readIOs, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.ReadIOs = readIOs
		case 4:
			readMerges, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.ReadMerges = readMerges
		case 5:
			readSectors, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.ReadSectors = readSectors
		case 6:
			readTicks, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.ReadTicks = readTicks
		case 7:
			writeIOs, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.WriteIOs = writeIOs
		case 8:
			writeMerges, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.WriteMerges = writeMerges
		case 9:
			writeSectors, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.WriteSectors = writeSectors
		case 10:
			writeTicks, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.WriteTicks = writeTicks
		case 11:
			inFlight, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.InFlight = inFlight
		case 12:
			ioTicks, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.IOTicks = ioTicks
		case 13:
			timeInQueue, _ := strconv.ParseUint(field, 10, 64)
			diskRawStats.TimeInQueue = timeInQueue
		}
	}

	return diskRawStats
}

func readProcDisksStats() []types.SysstatRaw {
	file, _ := os.Open("/proc/diskstats")
	stats := make([]types.SysstatRaw, 0)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	now := time.Now().Unix()

	for scanner.Scan() {
		out := parse(scanner.Text())
		out.SampleTime = now
		stats = append(stats, out)
	}

	return stats
}

func (c *DisksUtils) GetRawStats() []types.SysstatRaw {
	return readProcDisksStats()
}

func diskAvgStats(firstSample types.SysstatRaw, secondSample types.SysstatRaw) (diskAvgStats types.SysstatAvg) {
	diskAvgStats = types.SysstatAvg{}

	timeDelta := float64(secondSample.SampleTime - firstSample.SampleTime)

	// Check the samples are from the same disk
	if firstSample.Major != secondSample.Major ||
		firstSample.Minor != secondSample.Minor ||
		firstSample.Name != secondSample.Name {
		return types.SysstatAvg{}
	} else {
		diskAvgStats.Major = firstSample.Major
		diskAvgStats.Minor = firstSample.Minor
		diskAvgStats.Name = firstSample.Name
	}

	// Calculate average between the 2 samples
	diskAvgStats.ReadIOs = float64(secondSample.ReadIOs-firstSample.ReadIOs) / timeDelta
	diskAvgStats.ReadMerges = float64(secondSample.ReadMerges-firstSample.ReadMerges) / timeDelta
	diskAvgStats.ReadBytes = float64((secondSample.ReadSectors*512)-(firstSample.ReadSectors*512)) / timeDelta
	diskAvgStats.WriteIOs = float64(secondSample.WriteIOs-firstSample.WriteIOs) / timeDelta
	diskAvgStats.WriteMerges = float64(secondSample.WriteMerges-firstSample.WriteMerges) / timeDelta
	diskAvgStats.WriteBytes = float64((secondSample.WriteSectors*512)-(firstSample.WriteSectors*512)) / timeDelta

	diskAvgStats.InFlight = secondSample.InFlight
	diskAvgStats.TimeInQueue = secondSample.TimeInQueue - firstSample.TimeInQueue

	return diskAvgStats
}

func (c *DisksUtils) GetAvgStatsInterval(interval int64) []types.SysstatAvg {
	firstSamples := readProcDisksStats()

	diskAvgStatsArr := make([]types.SysstatAvg, 0)

	time.Sleep(time.Duration(interval) * time.Second)

	secondSamples := readProcDisksStats()

	for _, firstSample := range firstSamples {
		diskName := firstSample.Name
		matched, _ := regexp.Match("^sd\\w$|^nvme\\d\\w\\d$", []byte(diskName))

		if !matched {
			continue
		}

		for _, secondSample := range secondSamples {
			if secondSample.Name == diskName {
				diskAvg := diskAvgStats(firstSample, secondSample)

				diskAvgStatsArr = append(diskAvgStatsArr, diskAvg)
				break
			} else {
				continue
			}
		}
	}

	return diskAvgStatsArr
}
