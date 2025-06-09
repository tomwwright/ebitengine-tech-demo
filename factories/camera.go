package factories

import (
	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func CreateCamera(w donburi.World) *donburi.Entry {
	e := archetypes.Camera.Create(w)

	camera := components.Camera.Get(e)
	camera.Color = constants.White

	t := components.Transform.Get(e)
	scale := float64(constants.Scale)
	t.LocalPosition = math.NewVec2(0, 0)
	t.LocalScale = math.NewVec2(scale, scale)
	t.LocalRotation = 0

	return e
}
