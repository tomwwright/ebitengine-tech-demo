package events

import "github.com/yohamta/donburi/features/events"

type Interaction struct {
	Name string
}

var InteractionEvent = events.NewEventType[Interaction]()
