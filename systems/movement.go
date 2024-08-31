package systems

import (
	"techdemo/components"
	"techdemo/constants"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
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
			current, isFinished := movement.Tween.Update(constants.DeltaTime)
			movement.LastDirection = current.Sub(transform.LocalPosition).Normalized()
			transform.LocalPosition = current
			if isFinished {
				movement.Tween = nil
			}
		}
	})

}
