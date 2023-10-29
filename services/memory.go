package services

import "github.com/rafacas/sysstats"

type MemoryService struct {
}

func (ctrl *MemoryService) GetMemoryInfo() sysstats.MemStats {
	stat, _ := sysstats.GetMemStats()
	return stat
}
