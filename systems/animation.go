package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Animation struct {
	query *query.Query
}

func NewAnimation() *Animation {
	return &Animation{
		query: query.NewQuery(
			filter.Contains(components.Animation),
		),
	}
}

func (m *Animation) Update(ecs *ecs.ECS) {
	m.query.Each(ecs.World, func(entry *donburi.Entry) {
		animation := components.Animation.Get(entry)
		animation.Update(constants.DeltaTimeDuration)

		sprite := components.Sprite.Get(entry)
		if sprite != nil {
			sprite.Image = animation.Image()
		}
	})

}
