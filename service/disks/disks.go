package disks

import (
	"encoding/json"
	"os/exec"
	"time"
	"unraid-rest-api/service/disks/utils"
)

type DisksSysstats struct {
	history        []Stat
	count          int
	max            int
	parser         utils.DisksUtils
	arrayDirectory string
	mdStat         string
}

type Stat struct {
	Time time.Time
	Avg  []utils.DiskAvgStats
}

type DiskSysstatsRaw struct {
	Major        int    `json:"major"`        // Major number for the disk
	Minor        int    `json:"minor"`        // Minor number for the disk
	Name         string `json:"name"`         // Disk name
	ReadIOs      uint64 `json:"readios"`      // # of reads completed since boot
	ReadMerges   uint64 `json:"readmerges"`   // # of reads merged since boot
	ReadSectors  uint64 `json:"readsectors"`  // # of sectors read since boot
	ReadTicks    uint64 `json:"readticks"`    // # of milliseconds spent reading since boot
	WriteIOs     uint64 `json:"writeios"`     // # of writes completed since boot
	WriteMerges  uint64 `json:"writemerges"`  // # of writes merged since boot
	WriteSectors uint64 `json:"writesectors"` // # of sectors written since boot
	WriteTicks   uint64 `json:"writeticks"`   // # of milliseconds spent writing since boot
	InFlight     uint64 `json:"inflight"`     // # of I/Os currently in progress
	IOTicks      uint64 `json:"ioticks"`      // # of milliseconds spent doing I/Os since boot
	TimeInQueue  uint64 `json:"timeinqueue"`  // Weighted # of milliseconds spent doing I/Os since boot
	SampleTime   int64  `json:"sampletime"`   // Time when the sample was taken
}

type DiskInfo struct {
	Model string
}

type DiskLsblk struct {
	Blockdevices []struct {
		Alignment    int         `json:"alignment"`
		DiscAln      int         `json:"disc-aln"`
		Dax          bool        `json:"dax"`
		DiscGran     int         `json:"disc-gran"`
		DiscMax      int64       `json:"disc-max"`
		DiscZero     bool        `json:"disc-zero"`
		Fsavail      int64       `json:"fsavail"`
		Fsroots      []string    `json:"fsroots"`
		Fssize       int64       `json:"fssize"`
		Fstype       string      `json:"fstype"`
		Fsused       int64       `json:"fsused"`
		Fsuse        string      `json:"fsuse%"`
		Fsver        string      `json:"fsver"`
		Group        string      `json:"group"`
		Hctl         interface{} `json:"hctl"`
		Hotplug      bool        `json:"hotplug"`
		Kname        string      `json:"kname"`
		Label        interface{} `json:"label"`
		LogSec       int         `json:"log-sec"`
		MajMin       string      `json:"maj:min"`
		MinIo        int         `json:"min-io"`
		Mode         string      `json:"mode"`
		Model        interface{} `json:"model"`
		Name         string      `json:"name"`
		OptIo        int         `json:"opt-io"`
		Owner        string      `json:"owner"`
		Partflags    interface{} `json:"partflags"`
		Partlabel    interface{} `json:"partlabel"`
		Parttype     interface{} `json:"parttype"`
		Parttypename interface{} `json:"parttypename"`
		Partuuid     interface{} `json:"partuuid"`
		Path         string      `json:"path"`
		PhySec       int         `json:"phy-sec"`
		Pkname       interface{} `json:"pkname"`
		Pttype       interface{} `json:"pttype"`
		Ptuuid       interface{} `json:"ptuuid"`
		Ra           int         `json:"ra"`
		Rand         bool        `json:"rand"`
		Rev          interface{} `json:"rev"`
		Rm           bool        `json:"rm"`
		Ro           bool        `json:"ro"`
		Rota         bool        `json:"rota"`
		RqSize       int         `json:"rq-size"`
		Sched        string      `json:"sched"`
		Serial       interface{} `json:"serial"`
		Size         int         `json:"size"`
		Start        interface{} `json:"start"`
		State        interface{} `json:"state"`
		Subsystems   string      `json:"subsystems"`
		Mountpoint   string      `json:"mountpoint"`
		Mountpoints  []string    `json:"mountpoints"`
		Tran         interface{} `json:"tran"`
		Type         string      `json:"type"`
		UUID         interface{} `json:"uuid"`
		Vendor       interface{} `json:"vendor"`
		Wsame        int         `json:"wsame"`
		Wwn          interface{} `json:"wwn"`
		Zoned        string      `json:"zoned"`
		ZoneSz       int         `json:"zone-sz"`
		ZoneWgran    int         `json:"zone-wgran"`
		ZoneApp      int         `json:"zone-app"`
		ZoneNr       int         `json:"zone-nr"`
		ZoneOmax     int         `json:"zone-omax"`
		ZoneAmax     int         `json:"zone-amax"`
		Children     []struct {
			Alignment    int         `json:"alignment"`
			DiscAln      int         `json:"disc-aln"`
			Dax          bool        `json:"dax"`
			DiscGran     int         `json:"disc-gran"`
			DiscMax      int         `json:"disc-max"`
			DiscZero     bool        `json:"disc-zero"`
			Fsavail      int64       `json:"fsavail"`
			Fsroots      []string    `json:"fsroots"`
			Fssize       int64       `json:"fssize"`
			Fstype       string      `json:"fstype"`
			Fsused       int         `json:"fsused"`
			Fsuse        string      `json:"fsuse%"`
			Fsver        string      `json:"fsver"`
			Group        string      `json:"group"`
			Hctl         interface{} `json:"hctl"`
			Hotplug      bool        `json:"hotplug"`
			Kname        string      `json:"kname"`
			Label        string      `json:"label"`
			LogSec       int         `json:"log-sec"`
			MajMin       string      `json:"maj:min"`
			MinIo        int         `json:"min-io"`
			Mode         string      `json:"mode"`
			Model        interface{} `json:"model"`
			Name         string      `json:"name"`
			OptIo        int         `json:"opt-io"`
			Owner        string      `json:"owner"`
			Partflags    string      `json:"partflags"`
			Partlabel    interface{} `json:"partlabel"`
			Parttype     string      `json:"parttype"`
			Parttypename string      `json:"parttypename"`
			Partuuid     interface{} `json:"partuuid"`
			Path         string      `json:"path"`
			PhySec       int         `json:"phy-sec"`
			Pkname       string      `json:"pkname"`
			Pttype       string      `json:"pttype"`
			Ptuuid       interface{} `json:"ptuuid"`
			Ra           int         `json:"ra"`
			Rand         bool        `json:"rand"`
			Rev          interface{} `json:"rev"`
			Rm           bool        `json:"rm"`
			Ro           bool        `json:"ro"`
			Rota         bool        `json:"rota"`
			RqSize       int         `json:"rq-size"`
			Sched        string      `json:"sched"`
			Serial       interface{} `json:"serial"`
			Size         int64       `json:"size"`
			Start        int         `json:"start"`
			State        interface{} `json:"state"`
			Subsystems   string      `json:"subsystems"`
			Mountpoint   string      `json:"mountpoint"`
			Mountpoints  []string    `json:"mountpoints"`
			Tran         interface{} `json:"tran"`
			Type         string      `json:"type"`
			UUID         string      `json:"uuid"`
			Vendor       interface{} `json:"vendor"`
			Wsame        int         `json:"wsame"`
			Wwn          interface{} `json:"wwn"`
			Zoned        string      `json:"zoned"`
			ZoneSz       int         `json:"zone-sz"`
			ZoneWgran    int         `json:"zone-wgran"`
			ZoneApp      int         `json:"zone-app"`
			ZoneNr       int         `json:"zone-nr"`
			ZoneOmax     int         `json:"zone-omax"`
			ZoneAmax     int         `json:"zone-amax"`
		} `json:"children,omitempty"`
	} `json:"blockdevices"`
}

func NewDisksSysstats() *DisksSysstats {
	instance := &DisksSysstats{max: 120, mdStat: "/proc/mdstat"}
	instance.parser = utils.NewDiskUtils()

	go instance._go()

	return instance
}

func (s *DisksSysstats) _go() {
	for {

		avg := s.parser.GetAvgStatsInterval(1)
		wrapper := Stat{Avg: avg, Time: time.Now().UTC()}

		if s.count < s.max {
			s.history = append(s.history, wrapper)
			s.count++
		} else {
			s.history = append(s.history[1:], wrapper)
		}
	}
}

func (s *DisksSysstats) GetHistory() []Stat {
	return s.history
}

func (s *DisksSysstats) GetHistoryLast() Stat {
	return s.history[s.count-1]
}

func (s *DisksSysstats) GetDisksLsblk() DiskLsblk {

	output, _ := exec.Command("lsblk", "--json", "--output-all", "--bytes").Output()

	blockDevices := DiskLsblk{}

	json.Unmarshal(output, &blockDevices)

	return blockDevices
}
