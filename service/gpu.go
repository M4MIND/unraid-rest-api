package service

import (
	"unraid-rest-api/service/gpu"
	"unraid-rest-api/service/utils"
)

type GpuService struct {
	instance gpu.GpuInfoInstance
}

func NewGpuService() *GpuService {
	return &GpuService{
		instance: utils.NewNvidiaSmi(),
	}
}

func (c *GpuService) GetInfo() (stats gpu.GpuInfo, status bool) {
	return c.instance.GetInfo()
}
