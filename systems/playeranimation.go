package systems

import (
	"fmt"
	"techdemo/components"
	"techdemo/constants"
	"techdemo/tags"
	"techdemo/tilemap"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
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
	walkDown, _ := tilemap.GetAnimation("player", "walkDown")
	walkUp, _ := tilemap.GetAnimation("player", "walkUp")
	walkLeft, _ := tilemap.GetAnimation("player", "walkLeft")
	walkRight, _ := tilemap.GetAnimation("player", "walkRight")

	return &PlayerAnimation{
		query: donburi.NewQuery(filter.Contains(tags.Player, components.Movement, components.Animation)),
		animations: animations{
			idle:  idle,
			left:  walkLeft,
			right: walkRight,
			up:    walkUp,
			down:  walkDown,
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
			fmt.Printf("Setting animation %s\n", anim.Name)
			components.Animation.Set(playerEntry, &components.AnimationData{
				Durations: anim.Durations,
				Frames:    anim.Frames,
			})
			pa.currentAnimation = anim.Name
		}
	}

	movement := components.Movement.Get(playerEntry)
	animation := components.Animation.Get(playerEntry)

	isMoving := movement.Tween != nil
	if isMoving {
		animation.Resume()
	}

	direction := movement.LastDirection
	switch direction {

	case constants.Right:
		setAnimationComponent(pa.animations.right)

	case constants.Left:
		setAnimationComponent(pa.animations.left)

	case constants.Up:
		setAnimationComponent(pa.animations.up)

	case constants.Down:
		if isMoving {
			setAnimationComponent(pa.animations.down)
		} else {
			setAnimationComponent(pa.animations.idle)
		}

	default:
		setAnimationComponent(pa.animations.idle)

	}

	// idle animation is special case that should play while not moving
	if !isMoving && pa.currentAnimation != "idle" {
		animation.PauseAtStart()
	}
}
