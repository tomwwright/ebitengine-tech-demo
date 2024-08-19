package systems

import (
	"techdemo/components"
	"techdemo/tags"
	"techdemo/tilemap"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type PlayerAnimation struct {
	query            *query.Query
	currentAnimation string
	animations       animations
}

type animations struct {
	idle  tilemap.Animation
	left  tilemap.Animation
	right tilemap.Animation
	up    tilemap.Animation
	down  tilemap.Animation
}

func NewPlayerAnimation(tilemap *tilemap.Tilemap) *PlayerAnimation {
	idle, _ := tilemap.GetAnimation("player", "idle")
	walk, _ := tilemap.GetAnimation("player", "walk")
	walkUp, _ := tilemap.GetAnimation("player", "walkUp")

	return &PlayerAnimation{
		query: donburi.NewQuery(filter.Contains(tags.Player, components.Movement, components.Animation)),
		animations: animations{
			idle:  idle,
			left:  walk,
			right: walk,
			up:    walkUp,
			down:  walk,
		},
	}
}

func (pa *PlayerAnimation) Update(ecs *ecs.ECS) {

	playerEntry, ok := pa.query.First(ecs.World)
	if !ok {
		return
	}

	setAnimationComponent := func(anim tilemap.Animation) {
		if pa.currentAnimation != anim.Name {
			components.Animation.Set(playerEntry, &components.AnimationData{
				Durations: anim.Durations,
				Frames:    anim.Frames,
			})
			pa.currentAnimation = anim.Name
		}

	}

	movement := components.Movement.Get(playerEntry)
	if movement.Tween == nil {
		setAnimationComponent(pa.animations.idle)
		return
	}

	direction := movement.Tween.To.Sub(movement.Tween.From).Normalized()
	switch direction {

	case math.NewVec2(1.0, 0.0): // right
		setAnimationComponent(pa.animations.right)

	case math.NewVec2(-1.0, 0.0): // left
		setAnimationComponent(pa.animations.left)

	case math.NewVec2(0.0, -1.0): // up
		setAnimationComponent(pa.animations.up)

	case math.NewVec2(0.0, 1.0): // down
		setAnimationComponent(pa.animations.down)

	}
}
