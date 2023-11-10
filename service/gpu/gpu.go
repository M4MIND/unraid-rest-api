package gpu

import (
	"unraid-rest-api/service/gpu/types"
	"unraid-rest-api/service/gpu/utils"
)

type GpuService struct {
	instance types.GpuInfoInstance
}

func NewGpuService() *GpuService {
	return &GpuService{
		instance: utils.NewNvidiaSmi(),
	}
}

func (c *GpuService) GetInfo() (stats types.GpuInfo, status bool) {
	return c.instance.GetInfo()
}
