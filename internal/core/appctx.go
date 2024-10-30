package core

import (
	"time"
)

type AppContext struct {
	WorkingDir    string
	StartTime     time.Time
}
