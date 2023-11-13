package types

import (
	"time"
)

type Stat struct {
	Time time.Time
	Avg  []SysstatAvg
}
