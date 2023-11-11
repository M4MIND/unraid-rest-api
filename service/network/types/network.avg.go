package types

import (
	"github.com/rafacas/sysstats"
	"time"
)

type NetworkAvg struct {
	Avg  sysstats.NetAvgStats
	Time time.Time
}
