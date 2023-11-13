package http

import (
	"unraid-rest-api/api/disks"
	"unraid-rest-api/api/disks/types"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type disksHandler struct {
	services service.Container
}

func (d disksHandler) GetArrayInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		mdstats := d.services.RaidService.GetMdcmdStat()
		lsblk := d.services.DiskService.GetDisksLsblk()

		out := types.ArrayInfo{
			Size: 0,
			Used: 0,
			Free: 0,
		}

		for _, mdstatItem := range mdstats.Stats {
			smart := d.services.SmartService.GetDiskSmartInfo(mdstatItem.RdevName)

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

		context.JSON(200, out)
	}
}

func (d disksHandler) GetDisksInfoLsblk() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DiskService.GetDisksLsblk())
	}
}

func (d disksHandler) GetHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DiskService.GetHistoryLast())
	}
}

func (d disksHandler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DiskService.GetHistory())
	}
}

func NewHandler(s service.Container) disks.Handlers {
	return &disksHandler{
		services: s,
	}
}
