package disks

import (
	"encoding/json"
	"os/exec"
	"time"
	"unraid-rest-api/service/disks/types"
	"unraid-rest-api/service/disks/utils"
)

type Service struct {
	history        []types.Stat
	count          int
	max            int
	parser         utils.DisksUtils
	arrayDirectory string
	mdStat         string
}

func NewService() *Service {
	instance := &Service{max: 120, mdStat: "/proc/mdstat"}
	instance.parser = utils.NewDiskUtils()

	go instance._go()

	return instance
}

func (s *Service) _go() {
	for {

		avg := s.parser.GetAvgStatsInterval(1)
		wrapper := types.Stat{Avg: avg, Time: time.Now().UTC()}

		if s.count < s.max {
			s.history = append(s.history, wrapper)
			s.count++
		} else {
			s.history = append(s.history[1:], wrapper)
		}
	}
}

func (s *Service) GetHistory() []types.Stat {
	return s.history
}

func (s *Service) GetHistoryLast() types.Stat {
	return s.history[s.count-1]
}

func (s *Service) GetDisksLsblk() types.Lsblk {

	output, _ := exec.Command("lsblk", "--json", "--output-all", "--bytes").Output()

	blockDevices := types.Lsblk{}

	json.Unmarshal(output, &blockDevices)

	return blockDevices
}
