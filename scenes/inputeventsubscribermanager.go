package scenes

import (
	"techdemo/events"

	"github.com/yohamta/donburi"
	donburievents "github.com/yohamta/donburi/features/events"
)

type InputEventsSubscriber = donburievents.Subscriber[events.Input]

type InputEventsSubscriberManager struct {
	Subscribers              []InputEventsSubscriber
	DialogueStateSubscribers []InputEventsSubscriber
}

func (m *InputEventsSubscriberManager) OnStateChange(w donburi.World, change events.StateChange) {
	switch change {
	case events.DialogueClosed:
		m.OnDialogueClosed(w)
	case events.DialogueOpened:
		m.OnDialogueOpened(w)
	}
}

func (m *InputEventsSubscriberManager) OnDialogueOpened(w donburi.World) {
	for _, s := range m.Subscribers {
		events.InputEvent.Unsubscribe(w, s)
	}

	for _, s := range m.DialogueStateSubscribers {
		events.InputEvent.Subscribe(w, s)
	}
}

func (m *InputEventsSubscriberManager) OnDialogueClosed(w donburi.World) {
	for _, s := range m.Subscribers {
		events.InputEvent.Subscribe(w, s)
	}

	for _, s := range m.DialogueStateSubscribers {
		events.InputEvent.Unsubscribe(w, s)
	}
}
