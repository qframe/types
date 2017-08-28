package qtypes_interfaces

import (
	"time"
)

type BaseMessage interface {
	StopProcessing(p BasePlugin) bool
	GetTime() time.Time
}
