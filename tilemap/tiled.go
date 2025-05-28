package tilemap

import (
	"techdemo/interactions/yaml"

	"github.com/yohamta/donburi/features/math"
)

type Tiled interface {
	Size() (int, int)
	Objects() []Object
	Tiles() []Tile
	Layers() []Layer
	Animations() []Animation
	Interactions() *yaml.Interactions
	ToPixels(position math.Vec2) math.Vec2
	ToPosition(index int) math.Vec2
}
