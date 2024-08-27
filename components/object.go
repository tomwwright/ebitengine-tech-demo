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

func NewObject(offset math.Vec2, w float64, h float64, tags ...string) *ObjectData {
	return &ObjectData{
		Object:          *resolv.NewObject(0, 0, w, h),
		TransformOffset: offset,
	}
}
