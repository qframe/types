package qtypes_qframe

import (
	"github.com/qframe/types/plugin"
	"sync"
)

type QFrame struct {
	wg sync.WaitGroup
	registry *Registry
}

func NewQFrame() QFrame {
	return QFrame{
		registry: NewRegistry(),
	}
}

func (q *QFrame) AddPlugin(p *qtypes_plugin.Plugin) {
	q.registry.AddPlugin(p)
	q.wg.Add(1)
}