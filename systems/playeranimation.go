package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/tomwwright/ebitengine-tech-demo/tilemap"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var QueryPlayerAnimation = donburi.NewQuery(filter.Contains(tags.Player, components.Movement, components.Animation, components.CharacterAnimations))

const (
	AnimationKeyWalkRight = "player/walkRight"
	AnimationKeyWalkLeft  = "player/walkLeft"
	AnimationKeyWalkUp    = "player/walkUp"
	AnimationKeyWalkDown  = "player/walkDown"
	AnimationKeyIdle      = "player/idle"
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
			components.Animation.Set(playerEntry, components.NewAnimationFromTilemapAnimation(anim))
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
