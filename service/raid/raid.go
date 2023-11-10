package raid

import (
	"os"
	"regexp"
	"strconv"
	"unraid-rest-api/service/raid/types"
)

type RaidService struct {
}

func NewRaidService() RaidService {
	return RaidService{}
}

func (s *RaidService) GetMdcmdStat() types.Mdstats {
	mdstat, _ := os.ReadFile("/proc/mdstat")

	out := types.Mdstats{}

	convert := string(mdstat)

	disksNumber := regexp.MustCompile("(rdevNumber.\\d+)=(.+)").FindAllStringSubmatch(convert, -1)
	disksNames := regexp.MustCompile("(diskName.\\d+)=(.+)").FindAllStringSubmatch(convert, -1)
	disksSize := regexp.MustCompile("(diskSize.\\d+)=(\\d+)").FindAllStringSubmatch(convert, -1)
	disksState := regexp.MustCompile("(diskState.\\d+)=(\\d+)").FindAllStringSubmatch(convert, -1)
	disksId := regexp.MustCompile("(diskId.\\d+)=(.+)").FindAllStringSubmatch(convert, -1)
	diskRdevName := regexp.MustCompile("(rdevName.\\d+)=(.+)").FindAllStringSubmatch(convert, -1)
	diskRdevOffset := regexp.MustCompile("(rdevOffset.\\d+)=(.+)").FindAllStringSubmatch(convert, -1)

	for i := range disksNames {
		instance := types.Mdstat{}
		instance.DiskNumber, _ = strconv.Atoi(disksNumber[i][2])
		instance.DiskName = disksNames[i][2]
		instance.DiskSize, _ = strconv.Atoi(disksSize[i][2])
		instance.DiskState, _ = strconv.Atoi(disksState[i][2])
		instance.DiskId = disksId[i][2]
		instance.RdevName = diskRdevName[i][2]
		instance.RdevOffset, _ = strconv.Atoi(diskRdevOffset[i][2])

		out.Stats = append(out.Stats, instance)
	}

	return out
}
