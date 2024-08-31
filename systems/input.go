package systems

import (
	"techdemo/events"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/samber/lo"
	"github.com/yohamta/donburi/ecs"
)

type Input struct {
	activeInputs []events.Input
}

func NewInput() *Input {
	return &Input{}
}

func (input *Input) Update(ecs *ecs.ECS) {
	w := ecs.World
	pressed, active, released := getInputs()

	for _, e := range pressed {
		events.InputEvent.Publish(w, e)
	}

	if isMovementReleased(active, released) {
		events.InputEvent.Publish(w, events.InputMoveNone)
	}
}

// isMovementReleased returns true when active contains no movement inputs
// and released contains at least one movement input
func isMovementReleased(active []events.Input, released []events.Input) bool {
	activeMovements := lo.Intersect(movementInputs, active)
	releasedMovements := lo.Intersect(movementInputs, released)
	return len(activeMovements) == 0 && len(releasedMovements) > 0
}

func getInputs() (pressed []events.Input, active []events.Input, released []events.Input) {

	for input, key := range inputToKeyMapping {
		if inpututil.IsKeyJustPressed(key) {
			pressed = append(pressed, input)
		} else if ebiten.IsKeyPressed(key) {
			active = append(active, input)
		} else if inpututil.IsKeyJustReleased(key) {
			released = append(released, input)
		}
	}

	return pressed, active, released
}

var movementInputs = []events.Input{events.InputMoveUp, events.InputMoveDown, events.InputMoveLeft, events.InputMoveRight}

var inputToKeyMapping = map[events.Input]ebiten.Key{
	events.InputMoveUp:    ebiten.KeyUp,
	events.InputMoveDown:  ebiten.KeyDown,
	events.InputMoveLeft:  ebiten.KeyLeft,
	events.InputMoveRight: ebiten.KeyRight,
}
