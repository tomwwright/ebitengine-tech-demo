package events

import (
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
)

var MovementFinishedEvent = events.NewEventType[*donburi.Entry]()
