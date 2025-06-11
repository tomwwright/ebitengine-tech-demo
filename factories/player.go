package factories

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/components/collision"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/systems"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/tomwwright/ebitengine-tech-demo/tilemap"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type Animations interface {
	GetAnimation(name string) tilemap.Animation
}

func CreatePlayer(w donburi.World, animations Animations, position math.Vec2, layer int) (*donburi.Entry, error) {
	e := archetypes.Player.Create(w)
	transform := components.Transform.Get(e)

	transform.LocalPosition = position
	scale := float64(1)
	transform.LocalScale = math.NewVec2(scale, scale)

	object := components.NewObject(e, constants.TileSize, constants.TileSize, tags.ResolvTagInteractive)
	components.Object.Set(e, object)

	sprite := components.Sprite.Get(e)
	sprite.Layer = layer

	AddCollision(e, collision.CollisionBottom)

	ca := components.CharacterAnimations.Get(e)
	keys := []string{systems.AnimationKeyIdle, systems.AnimationKeyWalkUp, systems.AnimationKeyWalkDown, systems.AnimationKeyWalkRight, systems.AnimationKeyWalkLeft}
	for _, k := range keys {
		a := animations.GetAnimation(k)
		if a.Frames == nil {
			return nil, fmt.Errorf("unable to locate player animation: %s", k)
		}
		ca.Add(a)
	}

	return e, nil
}
