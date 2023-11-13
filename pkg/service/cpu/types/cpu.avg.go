package types

import (
	"github.com/rafacas/sysstats"
	"time"
)

type CpuAvg struct {
	Avg  sysstats.CpusAvgStats `json:"average"`
	Time time.Time             `json:"time"`
}
