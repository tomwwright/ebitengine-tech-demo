package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var zoomChangeQuery = donburi.NewQuery(filter.Contains(tags.ZoomChange, components.Tween, components.Target))

func UpdateZoomChange(ecs *ecs.ECS) {
	zoomChangeQuery.Each(ecs.World, func(e *donburi.Entry) {
		tween := components.Tween.Get(e)
		target := components.Target.GetValue(e)
		t := components.Transform.Get(target)

		zoom, isFinished := tween.Update(constants.DeltaTime)
		t.LocalScale.X = float64(zoom)
		t.LocalScale.Y = float64(zoom)

		if isFinished {
			e.Remove()
		}
	})
}
