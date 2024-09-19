package events

import "github.com/yohamta/donburi/features/events"

type StateChange string

const (
	DialogueOpened StateChange = "DialogueOpened"
	DialogueClosed StateChange = "DialogueClosed"
)

var StateChangeEvent = events.NewEventType[StateChange]()
