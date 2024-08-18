package systems

import (
	"techdemo/components"
	"techdemo/tags"
	"techdemo/tween"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Input struct {
	query *query.Query
}

func NewInput() *Input {
	return &Input{
		query: donburi.NewQuery(filter.Contains(tags.Player, transform.Transform, components.Movement)),
	}
}

func (input *Input) Update(ecs *ecs.ECS) {
	playerEntry, ok := input.query.First(ecs.World)
	if !ok {
		return
	}

	movement := components.Movement.Get(playerEntry)
	if movement.Tween != nil {
		return
	}

	direction := getInputDirection()
	if direction == inputNone {
		return
	}

	transform := components.Transform.Get(playerEntry)
	d := float32(8) * float32(transform.LocalScale.X)
	from := [2]float32{float32(transform.LocalPosition.X), float32(transform.LocalPosition.Y)}
	to := getMovement(from, direction, d)
	movement.Tween = tween.NewSliceTween(from[:], to[:], 0.12, ease.Linear)
}

type InputDirection int

const (
	inputNone InputDirection = iota
	inputDirectionUp
	inputDirectionDown
	inputDirectionLeft
	inputDirectionRight
)

func getInputDirection() InputDirection {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		return inputDirectionRight
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		return inputDirectionLeft
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		return inputDirectionUp
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		return inputDirectionDown
	}
	return inputNone
}

func getMovement(from [2]float32, direction InputDirection, d float32) [2]float32 {
	switch direction {
	case inputDirectionUp:
		return [2]float32{from[0], from[1] - d}
	case inputDirectionDown:
		return [2]float32{from[0], from[1] + d}
	case inputDirectionLeft:
		return [2]float32{from[0] - d, from[1]}
	case inputDirectionRight:
		return [2]float32{from[0] + d, from[1]}
	}
	return [2]float32{0, 0}
}
