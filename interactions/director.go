package interactions

import (
	"techdemo/events"
	"techdemo/interactions/internal/steps"
	"techdemo/sequences"

	"github.com/yohamta/donburi"
)

type Director struct {
	RunnableManager sequences.RunnableManager
}

func NewDirector() *Director {
	return &Director{
		RunnableManager: sequences.RunnableManager{},
	}
}

func (d *Director) OnInteractionEvent(w donburi.World, event events.Interaction) {
	sequence := &sequences.Sequence{
		Steps: []sequences.Runnable{
			&steps.DebugStep{
				Text: event.Name,
			},
			&steps.DialogueStep{
				Text:  event.Name,
				World: w,
			},
			&steps.DialogueStep{
				Text:  "More lembas bread?",
				World: w,
			},
		},
	}

	d.RunnableManager.Start(sequence)
}
