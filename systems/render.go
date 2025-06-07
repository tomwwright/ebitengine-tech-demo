package systems

import (
	"slices"
	"sort"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/samber/lo"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Render struct {
	sprites *query.Query
	texts   *query.Query
	cameras *query.Query
	buffer  *ebiten.Image
}

func NewRender() *Render {
	return &Render{
		sprites: query.NewQuery(
			filter.Contains(transform.Transform, components.Sprite),
		),
		texts: query.NewQuery(
			filter.Contains(transform.Transform, components.Text),
		),
		cameras: query.NewQuery(filter.Contains(transform.Transform, components.Camera)),
		buffer:  ebiten.NewImage(constants.ScreenWidth, constants.ScreenHeight),
	}
}

func (r *Render) Update(ecs *ecs.ECS) {
	r.cameras.Each(ecs.World, func(e *donburi.Entry) {
		position := transform.WorldPosition(e)
		rotation := transform.WorldRotation(e)
		scale := transform.WorldScale(e)
		camera := components.Camera.Get(e)
		camera.Calculate(position, scale, rotation)
	})
}

func (r *Render) Draw(ecs *ecs.ECS, screen *ebiten.Image) {

	e, ok := r.cameras.First(ecs.World)
	if !ok {
		return
	}
	camera := components.Camera.Get(e)
	camera.SetViewportFromImage(screen)

	entries := []*donburi.Entry{}
	r.sprites.Each(ecs.World, func(entry *donburi.Entry) {
		if isCullable(entry, camera) {
			return
		}

		entries = append(entries, entry)
	})

	sortEntriesForRendering(entries)

	byLayer := lo.GroupBy(entries, func(entry *donburi.Entry) int {
		return components.Sprite.Get(entry).Layer
	})
	layers := lo.Keys(byLayer)
	sort.Ints(layers)

	r.buffer.Clear()

	for _, layer := range layers {
		for _, entry := range byLayer[layer] {
			sprite := components.Sprite.Get(entry)
			if sprite.Image != nil {
				scale := transform.WorldScale(entry)
				position := transform.WorldPosition(entry)
				camera.Draw(sprite.Image, position, scale, r.buffer)
			}
		}
	}

	r.texts.Each(ecs.World, func(entry *donburi.Entry) {
		t := components.Text.Get(entry)
		scale := transform.WorldScale(entry)
		position := transform.WorldPosition(entry)
		camera.DrawText(t.Font, t.Text, position, scale, r.buffer)
	})

	var screenScaling = math.NewVec2(float64(screen.Bounds().Dx())/float64(constants.ScreenWidth), float64(screen.Bounds().Dy())/float64(constants.ScreenHeight))
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(screenScaling.X, screenScaling.Y)
	screen.DrawImage(r.buffer, op)
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

func isCullable(entry *donburi.Entry, camera *components.CameraData) bool {
	sprite := components.Sprite.Get(entry)
	if sprite.Image == nil {
		return true
	}
	scale := transform.WorldScale(entry)
	position := transform.WorldPosition(entry)
	size := position.Add(math.NewVec2(float64(sprite.Image.Bounds().Dx()), float64(sprite.Image.Bounds().Dy())).Mul(scale))

	return !camera.IsVisible(position, size)
}
