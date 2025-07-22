package events

import (
	"time"

	"github.com/yohamta/donburi/features/events"
)

var UpdateEvent = events.NewEventType[time.Duration]()
