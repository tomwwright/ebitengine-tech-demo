package events

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
)

var TriggerEvent = events.NewEventType[*donburi.Entry]()
