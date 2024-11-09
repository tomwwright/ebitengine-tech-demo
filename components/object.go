package components

import (
	"techdemo/constants"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type ObjectData struct {
	resolv.Object
	TransformOffset math.Vec2
}

var Object = donburi.NewComponentType[ObjectData]()

func NewObject(entry *donburi.Entry, collision CollisionType, tags ...string) *ObjectData {
	offset, w, h := getObjectCollision(collision)
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

type CollisionType string

const (
	CollisionFull   CollisionType = "full"
	CollisionNone   CollisionType = "none"
	CollisionLeft   CollisionType = "left"
	CollisionRight  CollisionType = "right"
	CollisionTop    CollisionType = "top"
	CollisionBottom CollisionType = "bottom"
)

func getObjectCollision(collision CollisionType) (offset math.Vec2, w float64, h float64) {
	half := float64(constants.TileSize / 2)
	switch collision {
	case CollisionLeft:
		offset = constants.Zero
		w = half
		h = constants.TileSize
	case CollisionRight:
		offset = math.NewVec2(half, 0)
		w = half
		h = constants.TileSize
	case CollisionTop:
		offset = constants.Zero
		w = constants.TileSize
		h = half
	case CollisionBottom:
		offset = math.NewVec2(0, half)
		w = constants.TileSize
		h = half
	default:
		offset = constants.Zero
		w = constants.TileSize
		h = constants.TileSize
	}

	return offset, w, h
}
