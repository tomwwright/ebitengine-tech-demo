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
	inputs := getInputs()

	// publish movement events whenever pressed
	for _, e := range movementInputs {
		if inputs[e] == statePressed {
			events.InputEvent.Publish(w, e)
		}
	}

	// publish move none input when all movements are released
	if isMovementReleased(inputs) {
		events.InputEvent.Publish(w, events.InputMoveNone)
	}

	// publish interact input only when no movements are active
	if !isMovementActive(inputs) && inputs[events.InputInteract] == statePressed {
		events.InputEvent.Publish(w, events.InputInteract)
	}

}

// isMovementReleased returns true when active contains no movement inputs
// and released contains at least one movement input
func isMovementReleased(inputs map[events.Input]inputState) bool {
	isMovementPressed := isMovementActive(inputs)
	releasedInputs := lo.Keys(lo.PickBy(inputs, func(input events.Input, state inputState) bool { return state == stateReleased }))
	releasedMovements := lo.Intersect(movementInputs, releasedInputs)
	return !isMovementPressed && len(releasedMovements) > 0
}

func isMovementActive(inputs map[events.Input]inputState) bool {
	activeInputs := lo.Keys(lo.PickBy(inputs, func(input events.Input, state inputState) bool { return state == stateActive }))
	activeMovements := lo.Intersect(movementInputs, activeInputs)
	return len(activeMovements) != 0
}

type inputState int

const (
	stateNone inputState = iota
	statePressed
	stateActive
	stateReleased
)

func getInputs() map[events.Input]inputState {
	inputs := map[events.Input]inputState{}
	for input, key := range inputToKeyMapping {
		if inpututil.IsKeyJustPressed(key) {
			inputs[input] = statePressed
		} else if ebiten.IsKeyPressed(key) {
			inputs[input] = stateActive
		} else if inpututil.IsKeyJustReleased(key) {
			inputs[input] = stateReleased
		}
	}

	return inputs
}

var movementInputs = []events.Input{events.InputMoveUp, events.InputMoveDown, events.InputMoveLeft, events.InputMoveRight}

var inputToKeyMapping = map[events.Input]ebiten.Key{
	events.InputMoveUp:    ebiten.KeyUp,
	events.InputMoveDown:  ebiten.KeyDown,
	events.InputMoveLeft:  ebiten.KeyLeft,
	events.InputMoveRight: ebiten.KeyRight,
	events.InputInteract:  ebiten.KeySpace,
}
