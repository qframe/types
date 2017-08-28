package qtypes_interfaces


type QInventory interface {
	SetItem(id string, val interface{})
	GetItem(id string) (val interface{}, err error)
}