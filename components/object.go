package components

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

var Object = donburi.NewComponentType[resolv.Object]()

func NewObject(entry *donburi.Entry, width int, height int, tags ...string) *resolv.Object {
	object := resolv.NewObject(0, 0, float64(width), float64(height), tags...)
	object.Data = entry

	return object
}

func ResolveObjectEntry(object *resolv.Object) *donburi.Entry {
	return object.Data.(*donburi.Entry)
}
