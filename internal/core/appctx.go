package core

import (
	"time"
)

// AppContext contains the context for the application - purely the "context" in which the application is running.
// This may be expanded, but should not include "calculated" values - just the raw context.
type AppContext struct {
	WorkingDir string
	StartTime  time.Time
}
