package events

import "github.com/yohamta/donburi/features/events"

type Input string

const (
	InputMoveNone  Input = "InputMoveNone"
	InputMoveUp    Input = "InputMoveUp"
	InputMoveDown  Input = "InputMoveDown"
	InputMoveLeft  Input = "InputMoveLeft"
	InputMoveRight Input = "InputMoveRight"
	InputInteract  Input = "InputInteract"
)

var InputEvent = events.NewEventType[Input]()
