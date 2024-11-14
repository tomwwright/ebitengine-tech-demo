package steps

import (
	"techdemo/events"
	"techdemo/factories/dialogue"
	"techdemo/sequences"

	"github.com/yohamta/donburi"
)

type DialogueStep struct {
	World donburi.World
	Text  string
	done  sequences.Done
}

func (ds *DialogueStep) Run(done sequences.Done) {
	ds.done = done

	dialogue.CreateDialogue(ds.World, ds.Text)
	events.StateChangeEvent.Subscribe(ds.World, ds.onDialogueClosed)
}

func (ds *DialogueStep) onDialogueClosed(w donburi.World, event events.StateChange) {
	if event == events.DialogueClosed {
		events.StateChangeEvent.Unsubscribe(w, ds.onDialogueClosed)
		ds.done()
	}
}
