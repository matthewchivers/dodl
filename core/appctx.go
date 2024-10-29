package core

import (
	"time"
)

type AppContext struct {
	WorkspaceRoot string
	WorkingDir    string
	StartTime     time.Time
}
