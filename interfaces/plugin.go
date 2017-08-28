package qtypes_interfaces

import (
	"github.com/qframe/types/qchannel"
	"github.com/zpatrick/go-config"
)

// QPlugin
type BasePlugin interface {
	// Starts consumes bcast channels, config and attributes of a Plugin and runs a loop
	GetInfo() (typ,pkg,name string)
	CfgStringOr(path, alt string) string
	GetLogOnlyPlugs() []string
	GetChannels() qtypes_qchannel.QChan
	GetConfig() *config.Config
	SendData()
}

type QPlugin interface {
	Run()
}