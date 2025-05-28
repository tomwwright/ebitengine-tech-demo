package factories

import (
	"fmt"
	"techdemo/components"
	"techdemo/components/collision"
	"techdemo/systems"
	"techdemo/tags"
	"techdemo/tilemap"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type Animations interface {
	GetAnimation(name string) tilemap.Animation
}

func CreatePlayer(w donburi.World, animations Animations, position math.Vec2, layer int) (*donburi.Entry, error) {

	entity := w.Create(tags.Player, components.Transform, components.Sprite, components.Object, components.Movement, components.Animation, components.CharacterAnimations)
	entry := w.Entry(entity)

	transform := components.Transform.Get(entry)

	transform.LocalPosition = position
	scale := float64(1)
	transform.LocalScale = math.NewVec2(scale, scale)

	sprite := components.Sprite.Get(entry)
	sprite.Layer = layer

	object := components.NewObject(entry, collision.CollisionBottom, tags.ResolvTagCollider)
	components.Object.Set(entry, object)

	ca := components.CharacterAnimations.Get(entry)
	keys := []string{systems.AnimationKeyIdle, systems.AnimationKeyWalkUp, systems.AnimationKeyWalkDown, systems.AnimationKeyWalkRight, systems.AnimationKeyWalkLeft}
	for _, k := range keys {
		a := animations.GetAnimation(k)
		if a.Frames == nil {
			return nil, fmt.Errorf("unable to locate player animation: %s", k)
		}
		ca.Add(a)
	}

	return entry, nil
}
