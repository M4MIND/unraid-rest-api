package types

type Lsblk struct {
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
		Size         int64       `json:"size"`
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
