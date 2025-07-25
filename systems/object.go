package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

var updateObjectsQuery = donburi.NewQuery(filter.Contains(components.Transform, components.Object))

type ObjectsSystem struct {
	Space *resolv.Space
}

func NewObjects() *ObjectsSystem {
	return &ObjectsSystem{}
}

func (os *ObjectsSystem) Update(ecs *ecs.ECS) {
	updateObjectsQuery.Each(ecs.World, func(e *donburi.Entry) {
		object := components.Object.Get(e)

		if os.Space != nil && object.Space != os.Space {
			os.Space.Add(object)
		}

		position := transform.WorldPosition(e)
		object.Position.X = position.X
		object.Position.Y = position.Y
		object.Update()
	})
}
