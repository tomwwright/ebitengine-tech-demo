package steps

import (
	"fmt"

	"github.com/tanema/gween/ease"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/interactions/yaml"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/tomwwright/ebitengine-tech-demo/tween"

	"github.com/yohamta/donburi"
)

type MovementStep struct {
	yaml.Movement
	World donburi.World
}

func (s *MovementStep) Run(done sequences.Done) {
	defer done()

	v := constants.Zero
	switch s.Direction {
	case "up":
		v = constants.Up
	case "down":
		v = constants.Down
	case "left":
		v = constants.Left
	case "right":
		v = constants.Right
	default:
		panic(fmt.Errorf("invalid direction string: %s", s.Movement))
	}
	v = v.MulScalar(float64(constants.TileSize * s.Distance))

	e := tags.Player.MustFirst(s.World)
	movement := components.Movement.Get(e)
	transform := components.Transform.Get(e)

	if movement.Tween != nil {
		panic(fmt.Errorf("unable to set movement for already moving component: %+v %+v", e, movement))
	}

	from := transform.LocalPosition
	to := from.Add(v)
	movement.Tween = tween.NewVec2Tween(from, to, float32(constants.MovementSpeed)*float32(v.Magnitude()), ease.Linear)
}
