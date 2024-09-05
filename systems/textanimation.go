package systems

import (
	"techdemo/components"
	"techdemo/constants"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type TextAnimation struct {
	query *query.Query
}

func NewTextAnimation() *TextAnimation {
	return &TextAnimation{
		query: query.NewQuery(
			filter.Contains(components.Text, components.TextAnimation),
		),
	}
}

func (m *TextAnimation) Update(ecs *ecs.ECS) {
	m.query.Each(ecs.World, func(entry *donburi.Entry) {
		animation := components.TextAnimation.Get(entry)
		text := components.Text.Get(entry)
		if !animation.IsFinished() {
			animation.Characters += animation.Speed * constants.DeltaTime

			length := int(animation.Characters) - 1
			if length >= 0 {
				text.Text = animation.Text[:length]
			}
		}
	})

}
