package core

import (
	"time"
)

// AppContext contains the context for the application - purely the "context" in which the application is running.
// This may be expanded, but should not include "calculated" values - just the raw context.
type AppContext struct {
	// WorkingDir is the working directory of the application.
	WorkingDir string

	// StartTime is the time that the application started.
	StartTime time.Time

	// ReferenceTime is the time that is used as the reference point for the application.
	ReferenceTime time.Time
}
