package events

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
)

type Interaction struct {
	Name   string
	Target *donburi.Entry
}

var InteractionEvent = events.NewEventType[Interaction]()
var InteractionFinishedEvent = events.NewEventType[Interaction]()
