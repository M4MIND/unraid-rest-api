package ws

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/api/disks/types"
	"unraid-rest-api/api/websocket"
	"unraid-rest-api/pkg/service"
)

type Handler struct {
	service service.Container
}

func (h Handler) MemoryInfoTick() (interface{}, error) {
	return h.service.MemoryService.GetHistoryLast(), nil
}

func (h Handler) ArrayInfo() (interface{}, error) {
	mdstats := h.service.RaidService.GetMdcmdStat()
	lsblk := h.service.DiskService.GetDisksLsblk()

	out := types.ArrayInfo{
		Size: 0,
		Used: 0,
		Free: 0,
	}

	for _, mdstatItem := range mdstats.Stats {
		smart := h.service.SmartService.GetDiskSmartInfo(mdstatItem.RdevName)

		wrapper := types.ArrayInfoDevice{
			DiskId:      mdstatItem.DiskId,
			DiskState:   mdstatItem.DiskState,
			RdevName:    mdstatItem.RdevName,
			Temperature: smart.Temperature.Current,
		}

		for _, lsblkItem := range lsblk.Blockdevices {
			if lsblkItem.Name == wrapper.RdevName {
				wrapper.DiskSizeBytes = lsblkItem.Size
				wrapper.IsHdd = lsblkItem.Rota
				out.Size += lsblkItem.Size
				continue
			}
			if lsblkItem.Kname == mdstatItem.DiskName {
				wrapper.DiskUsedPercent = float32(lsblkItem.Fsused) / float32(lsblkItem.Size) * 100
				wrapper.DiskUsedBytes = lsblkItem.Fsused
				wrapper.Mount = lsblkItem.Mountpoint
				out.Used += lsblkItem.Fsused
				continue
			}
		}

		out.Devices = append(out.Devices, wrapper)
	}

	out.Free = out.Size - out.Used

	return out, nil
}

func (h Handler) PingPong() (interface{}, error) {
	return gin.H{"Ping": "Pong"}, nil
}

func (h Handler) CpuState() (interface{}, error) {
	return h.service.CpuService.GetAvgHistoryLast(), nil
}

func NewHandler(s service.Container) websocket.Handlers {
	return &Handler{service: s}
}
