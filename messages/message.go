package qtypes_messages

import (
	"encoding/json"
	"github.com/deckarep/golang-set"
	"github.com/qframe/types/plugin"
	"fmt"
)

type Message struct {
	Base
	Message string
}

func NewMessage(b Base, msg string) Message {
	m := Message{
		Base: b,
		Message: msg,
	}
	m.GenDefaultID()
	return m
}

func (m *Message) ParseJsonMap(p *qtypes_plugin.Plugin, keys mapset.Set, kv map[string]string) {
	it := keys.Iterator()
	for val := range it.C {
		key := val.(string)
		v, ok := kv[key]
		if !ok {
			p.Log("debug", fmt.Sprintf("Could not find key '%s' in Tags: %v", key, kv))
			continue
		}
		p.Log("debug", fmt.Sprintf("unmarshall: %s", v))
		byt := []byte(v)
		var dat map[string]interface{}
		json.Unmarshal(byt, &dat)
		for k, v := range dat {
			if _, ok := m.Tags[k]; !ok {
				p.Log("debug", fmt.Sprintf("New key in tag '%s' for message '%s'", k, v))
				m.Tags[k] = fmt.Sprintf("%s", v)
			} else {
				p.Log("debug", fmt.Sprintf("Overwrite tag '%s' in message '%s'", k, v))
				m.Tags[k] = fmt.Sprintf("%s", v)
			}
		}
	}

}