package qtypes_qframe

import (
	"fmt"
	"sync"
	"github.com/qframe/types/plugin"
)


type Plugins struct {
	state string
	plugin *qtypes_plugin.Plugin
}

type Registry struct {
	mu sync.Mutex
	Plugins map[string]Plugins
}


func NewRegistry() *Registry {
	return &Registry{
		Plugins: map[string]Plugins{},
	}
}

func (r *Registry) AddPlugin(p *qtypes_plugin.Plugin) {
	r.mu.Lock()
	defer r.mu.Unlock()
	key := genKey(p)
	r.Plugins[key] = Plugins{
		state: "added",
		plugin: p,
	}
}

func genKey(plugin *qtypes_plugin.Plugin) string {
	return fmt.Sprintf("%s.%s.%s", plugin.Typ, plugin.Pkg, plugin.Name)
}