package systems

import (
	"techdemo/components"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
)

var updateObjectsQuery = donburi.NewQuery(filter.Contains(components.Transform, components.Object))

func UpdateObjects(ecs *ecs.ECS) {
	updateObjectsQuery.Each(ecs.World, func(e *donburi.Entry) {
		object := components.Object.Get(e)
		position := transform.WorldPosition(e).Add(object.TransformOffset)
		object.Position.X = position.X
		object.Position.Y = position.Y
		object.Update()
	})
}
