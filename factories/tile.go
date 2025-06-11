package factories

import (
	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/components/collision"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func CreateTile(w donburi.World, position math.Vec2, layer int, collisionType collision.CollisionType, image *ebiten.Image) *donburi.Entry {
	e := archetypes.Sprite.Create(w)
	transform := components.Transform.Get(e)

	transform.LocalPosition = position
	scale := float64(1)
	transform.LocalScale = math.NewVec2(scale, scale)

	sprite := components.Sprite.Get(e)
	sprite.Image = image
	sprite.Layer = layer

	if collisionType != collision.CollisionNone {
		AddCollision(e, collisionType)
	}

	return e
}
