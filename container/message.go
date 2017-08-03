package container

/*
import (
	"fmt"
	"strings"
	"github.com/docker/docker/api/types"
	"github.com/qframe/types/messages"
)


func NewContainerMessage(base qtypes_messages.Base, cnt types.ContainerJSON, name, mType, msg string) qtypes_messages.Message {
	m := qtypes_messages.NewMessage(base, name, mType, msg)
	m.Container = cnt
	m.ID = m.GenContainerMsgID()
	return m
}

// GenContainerMsgID uses "<container_id>-<time.UnixNano()>-<MSG>" and does a sha1 hash.
func (m *Message) GenContainerMsgID() string {
	s := fmt.Sprintf("%s-%d-%s", m.Container.ID, m.Time.UnixNano(), m.Message)
	return Sha1HashString(s)
}

func (m *Message) GetContainerName() string {
	if m.Container.Name != "" {
		return strings.Trim(m.Container.Name, "/")
	} else {
		return "<none>"
	}
}
*/