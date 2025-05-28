package components

import (
	"github.com/tomwwright/ebitengine-tech-demo/components/collision"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type ObjectData struct {
	resolv.Object
	TransformOffset math.Vec2
}

var Object = donburi.NewComponentType[ObjectData]()

func NewObject(entry *donburi.Entry, collision collision.CollisionType, tags ...string) *ObjectData {
	offset, w, h := collision.Mask()
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
