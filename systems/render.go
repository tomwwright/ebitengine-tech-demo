package systems

import (
	"sort"
	"techdemo/components"
	"techdemo/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	vec2 "github.com/yohamta/donburi/features/math"
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

	var cameraMatrix ebiten.GeoM
	if camera, ok := tags.Camera.First(ecs.World); ok {
		viewport := vec2.NewVec2(float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy()))
		cameraMatrix = toCameraSpace(camera, viewport)
	}

	for _, layer := range layers {
		for _, entry := range byLayer[layer] {
			sprite := components.Sprite.Get(entry)
			if sprite.Image != nil {
				op := &ebiten.DrawImageOptions{}

				scale := transform.WorldScale(entry)
				op.GeoM.Scale(scale.X, scale.Y)

				position := transform.WorldPosition(entry)
				op.GeoM.Translate(position.X, position.Y)

				op.GeoM.Concat(cameraMatrix)

				screen.DrawImage(sprite.Image, op)
			}
		}
	}
}

func toCameraSpace(cameraEntry *donburi.Entry, viewport vec2.Vec2) ebiten.GeoM {
	m := ebiten.GeoM{}

	viewportCenter := viewport.MulScalar(0.5)
	position := transform.WorldPosition(cameraEntry)
	rotation := transform.WorldRotation(cameraEntry)
	scale := transform.WorldScale(cameraEntry)

	m.Translate(-position.X, -position.Y)

	m.Translate(-viewportCenter.X/scale.X, -viewportCenter.Y/scale.Y) // rotate around center of viewport, considering scale
	m.Rotate(rotation)

	m.Translate(viewportCenter.X/scale.X, viewportCenter.Y/scale.Y)

	m.Scale(
		scale.X,
		scale.Y,
	)

	return m
}
