package disks

import (
	"time"
	"unraid-rest-api/service/disks/types"
	"unraid-rest-api/service/disks/utils/lsblk"
	"unraid-rest-api/service/disks/utils/sysstat"
)

type Service struct {
	history        []types.Stat
	count          int
	max            int
	parser         sysstat.DisksUtils
	arrayDirectory string
	lsblk          *lsblk.Lsblk
}

func NewService() *Service {
	instance := &Service{max: 120, lsblk: lsblk.NewService()}
	instance.parser = sysstat.NewDiskUtils()

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
	return s.lsblk.GetInfo()
}
