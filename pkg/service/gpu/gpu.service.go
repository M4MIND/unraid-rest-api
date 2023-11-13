package gpu

import (
	"unraid-rest-api/pkg/service/gpu/types"
	"unraid-rest-api/pkg/service/gpu/utils"
)

type Service struct {
	instance types.GpuInfoInstance
}

func NewService() *Service {
	return &Service{
		instance: utils.NewNvidiaSmi(),
	}
}

func (c *Service) GetInfo() (stats types.GpuInfo, status bool) {
	return c.instance.GetInfo()
}
