package systems

import (
	"techdemo/components"
	"techdemo/constants"
	"techdemo/tags"
	"techdemo/tilemap"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var QueryPlayerAnimation = donburi.NewQuery(filter.Contains(tags.Player, components.Movement, components.Animation, components.CharacterAnimations))

const (
	AnimationKeyWalkRight = "walkRight"
	AnimationKeyWalkLeft  = "walkLeft"
	AnimationKeyWalkUp    = "walkUp"
	AnimationKeyWalkDown  = "walkDown"
	AnimationKeyIdle      = "idle"
)

func UpdatePlayerAnimation(ecs *ecs.ECS) {

	playerEntry, ok := QueryPlayerAnimation.First(ecs.World)
	if !ok {
		return
	}

	movement := components.Movement.Get(playerEntry)
	animation := components.Animation.Get(playerEntry)
	characterAnimations := components.CharacterAnimations.Get(playerEntry)

	setAnimationComponent := func(anim tilemap.Animation) {
		if animation.Name != anim.Name {
			components.Animation.Set(playerEntry, &components.AnimationData{
				Durations: anim.Durations,
				Frames:    anim.Frames,
				Name:      anim.Name,
			})
			animation = components.Animation.Get(playerEntry)
		}
	}

	isMoving := movement.Tween != nil
	if isMoving {
		animation.Resume()
	}

	direction := movement.LastDirection
	switch direction {

	case constants.Right:
		setAnimationComponent(characterAnimations.Animations[AnimationKeyWalkRight])

	case constants.Left:
		setAnimationComponent(characterAnimations.Animations[AnimationKeyWalkLeft])

	case constants.Up:
		setAnimationComponent(characterAnimations.Animations[AnimationKeyWalkUp])

	case constants.Down:
		if isMoving {
			setAnimationComponent(characterAnimations.Animations[AnimationKeyWalkDown])
		} else {
			setAnimationComponent(characterAnimations.Animations[AnimationKeyIdle])
		}

	default:
		setAnimationComponent(characterAnimations.Animations[AnimationKeyIdle])

	}

	// idle animation is special case that should play while not moving
	if !isMoving && animation.Name != AnimationKeyIdle {
		animation.PauseAtEnd()
	}
}
