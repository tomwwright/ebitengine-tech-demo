package factories

import (
	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

func CreateObject(w donburi.World, position math.Vec2, name string) *donburi.Entry {
	e := archetypes.Interaction.Create(w)

	transform := components.Transform.Get(e)
	transform.LocalPosition = position

	object := components.NewObject(e, constants.TileSize, constants.TileSize, tags.ResolvTagInteractive)
	components.Object.Set(e, object)

	interaction := components.Interaction.Get(e)
	interaction.Name = name

	return e
}
