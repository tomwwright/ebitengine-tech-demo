package systems

import (
	"techdemo/components"
	"techdemo/constants"
	"techdemo/events"
	"techdemo/tags"
	"techdemo/tween"

	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/donburi/query"
)

type Direction int

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

type PlayerMovement struct {
	query            *query.Query
	currentDirection Direction
}

func NewPlayerMovement() *PlayerMovement {
	return &PlayerMovement{
		query: donburi.NewQuery(filter.Contains(tags.Player, transform.Transform, components.Movement, components.Object)),
	}
}

func (m *PlayerMovement) OnInputEvent(w donburi.World, input events.Input) {
	m.currentDirection = toDirection(input)
}

func (m *PlayerMovement) Update(ecs *ecs.ECS) {
	playerEntry, ok := m.query.First(ecs.World)
	if !ok {
		return
	}

	movement := components.Movement.Get(playerEntry)
	if movement.Tween != nil {
		return
	}

	direction := m.currentDirection
	if direction == DirectionNone {
		return
	}

	d := float64(constants.TileSize / 2)
	v := toMovementVector(direction)
	delta := v.MulScalar(d)

	object := components.Object.Get(playerEntry)

	if collision := object.Check(delta.XY()); collision != nil {
		return
	}

	transform := components.Transform.Get(playerEntry)
	from := transform.LocalPosition
	to := from.Add(delta)
	movement.Tween = tween.NewVec2Tween(from, to, 0.12, ease.Linear)
}

func toDirection(input events.Input) Direction {
	switch input {
	case events.InputMoveUp:
		return DirectionUp
	case events.InputMoveDown:
		return DirectionDown
	case events.InputMoveLeft:
		return DirectionLeft
	case events.InputMoveRight:
		return DirectionRight
	default:
		return DirectionNone
	}
}

func toMovementVector(direction Direction) math.Vec2 {
	switch direction {
	case DirectionUp:
		return constants.Up
	case DirectionDown:
		return constants.Down
	case DirectionLeft:
		return constants.Left
	case DirectionRight:
		return constants.Right
	default:
		return constants.Zero
	}
}
