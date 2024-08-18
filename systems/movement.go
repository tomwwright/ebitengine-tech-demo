package systems

import (
	"techdemo/components"
	"techdemo/config"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Movement struct {
	query *query.Query
}

func NewMovement() *Movement {
	return &Movement{
		query: query.NewQuery(
			filter.Contains(transform.Transform, components.Movement),
		),
	}
}

func (m *Movement) Update(ecs *ecs.ECS) {
	m.query.Each(ecs.World, func(entry *donburi.Entry) {
		movement := components.Movement.Get(entry)
		if movement.Tween != nil {
			transform := components.Transform.Get(entry)
			current, isFinished := movement.Tween.Update(config.DeltaTime)
			transform.LocalPosition = math.NewVec2(float64(current[0]), float64(current[1]))
			if isFinished {
				movement.Tween = nil
			}
		}
	})

}
