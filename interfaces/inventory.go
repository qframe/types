package qtypes_interfaces

import (
	"github.com/moby/moby/api/types"
)

type ContainerInventory interface {
	SetItem(id string, cnt types.ContainerJSON)
	GetItem(id string) (cnt types.ContainerJSON, err error)
}