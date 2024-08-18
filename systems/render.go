package systems

import (
	"sort"
	"techdemo/components"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Render struct {
	query *query.Query
}

func NewRender() *Render {
	return &Render{
		query: query.NewQuery(
			filter.Contains(transform.Transform, components.Sprite),
		),
	}
}

func (r *Render) Update(ecs *ecs.ECS) {
	// do nothing
}

func (r *Render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	var entries []*donburi.Entry
	r.query.Each(ecs.World, func(entry *donburi.Entry) {
		entries = append(entries, entry)
	})

	byLayer := lo.GroupBy(entries, func(entry *donburi.Entry) int {
		return int(components.Sprite.Get(entry).Layer)
	})
	layers := lo.Keys(byLayer)
	sort.Ints(layers)

	for _, layer := range layers {
		for _, entry := range byLayer[layer] {
			sprite := components.Sprite.Get(entry)

			op := &ebiten.DrawImageOptions{}

			scale := transform.WorldScale(entry)
			op.GeoM.Scale(scale.X, scale.Y)

			position := transform.WorldPosition(entry)
			op.GeoM.Translate(position.X, position.Y)

			screen.DrawImage(sprite.Image, op)
		}
	}
}
