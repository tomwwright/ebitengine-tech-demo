package factories

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/components/collision"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func CreateTile(w donburi.World, position math.Vec2, layer int, collisionType collision.CollisionType, image *ebiten.Image) *donburi.Entry {
	entity := w.Create(components.Transform, components.Sprite)
	entry := w.Entry(entity)

	transform := components.Transform.Get(entry)

	transform.LocalPosition = position
	scale := float64(1)
	transform.LocalScale = math.NewVec2(scale, scale)

	sprite := components.Sprite.Get(entry)
	sprite.Image = image
	sprite.Layer = layer

	if collisionType != collision.CollisionNone {
		AddCollision(entry, collisionType)
	}

	return entry
}
