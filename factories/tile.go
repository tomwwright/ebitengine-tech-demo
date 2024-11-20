package factories

import (
	"techdemo/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func CreateTile(w donburi.World, t *tiled.LayerTile, position math.Vec2, layer int, collision components.CollisionType, image *ebiten.Image) *donburi.Entry {
	entity := w.Create(components.Transform, components.Sprite)
	entry := w.Entry(entity)

	transform := components.Transform.Get(entry)

	transform.LocalPosition = position
	scale := float64(1)
	transform.LocalScale = math.NewVec2(scale, scale)

	sprite := components.Sprite.Get(entry)
	sprite.Image = image
	sprite.Layer = layer

	if collision != components.CollisionNone {
		object := components.NewObject(entry, collision)
		entry.AddComponent(components.Object)
		components.Object.Set(entry, object)
	}

	return entry
}
