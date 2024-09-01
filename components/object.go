package components

import (
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type ObjectData struct {
	resolv.Object
	TransformOffset math.Vec2
}

var Object = donburi.NewComponentType[ObjectData]()

func NewObject(entry *donburi.Entry, offset math.Vec2, w float64, h float64, tags ...string) *ObjectData {
	object := resolv.NewObject(0, 0, w, h, tags...)
	object.Data = entry
	return &ObjectData{
		Object:          *object,
		TransformOffset: offset,
	}
}

func ResolveObjectEntry(object *resolv.Object) *donburi.Entry {
	return object.Data.(*donburi.Entry)
}
