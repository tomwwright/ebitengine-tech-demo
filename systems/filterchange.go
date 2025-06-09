package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var filterChangeQuery = donburi.NewQuery(filter.Contains(tags.FilterChange, components.TweenColor, components.Target))

func UpdateFilterChange(ecs *ecs.ECS) {
	filterChangeQuery.Each(ecs.World, func(e *donburi.Entry) {
		t := components.TweenColor.Get(e)
		target := components.Target.GetValue(e)
		camera := components.Camera.Get(target)

		color, isFinished := t.Update(constants.DeltaTime)
		camera.Color = color

		if isFinished {
			e.Remove()
		}
	})
}
