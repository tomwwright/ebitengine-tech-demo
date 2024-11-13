package systems

import (
	"slices"
	"sort"
	"techdemo/components"
	"techdemo/tags"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	vec2 "github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

const LineSpacing = 16

type Render struct {
	sprites *query.Query
	texts   *query.Query
}

func NewRender() *Render {
	return &Render{
		sprites: query.NewQuery(
			filter.Contains(transform.Transform, components.Sprite),
		),
		texts: query.NewQuery(
			filter.Contains(transform.Transform, components.Text),
		),
	}
}

func (r *Render) Update(ecs *ecs.ECS) {
	// do nothing
}

func (r *Render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {
	var entries []*donburi.Entry
	r.sprites.Each(ecs.World, func(entry *donburi.Entry) {
		entries = append(entries, entry)
	})

	sortEntriesForRendering(entries)

	byLayer := lo.GroupBy(entries, func(entry *donburi.Entry) int {
		return components.Sprite.Get(entry).Layer
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

	r.texts.Each(ecs.World, func(entry *donburi.Entry) {
		op := &text.DrawOptions{
			LayoutOptions: text.LayoutOptions{
				LineSpacing: LineSpacing,
			},
		}

		t := components.Text.Get(entry)

		scale := transform.WorldScale(entry)
		op.GeoM.Scale(scale.X, scale.Y)

		position := transform.WorldPosition(entry)
		op.GeoM.Translate(position.X, position.Y)

		op.GeoM.Concat(cameraMatrix)

		text.Draw(screen, t.Text, t.Font, op)
	})
}

func sortEntriesForRendering(entries []*donburi.Entry) {
	slices.SortFunc(entries, func(entryA *donburi.Entry, entryB *donburi.Entry) int {
		a := transform.WorldPosition(entryA)
		b := transform.WorldPosition(entryB)

		diff := int(a.Y - b.Y)

		if diff != 0 {
			return diff
		}

		return sublayerOrder(entryA) - sublayerOrder(entryB)
	})
}

// sublayerOrder returns an int for comparing render order
// of overlapping objects within a layer
func sublayerOrder(entry *donburi.Entry) int {
	// things that don't move sit above things that do
	if entry.HasComponent(components.Movement) {
		return 0
	} else {
		return 1
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
