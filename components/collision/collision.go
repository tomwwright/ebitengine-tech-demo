package collision

import (
	"techdemo/constants"

	"github.com/yohamta/donburi/features/math"
)

type CollisionType string

const (
	CollisionUndefined CollisionType = ""
	CollisionFull      CollisionType = "full"
	CollisionNone      CollisionType = "none"
	CollisionLeft      CollisionType = "left"
	CollisionRight     CollisionType = "right"
	CollisionTop       CollisionType = "top"
	CollisionBottom    CollisionType = "bottom"
)

func (collision CollisionType) Mask() (offset math.Vec2, w float64, h float64) {
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
