package systems

import (
	"techdemo/components"
	"techdemo/events"
	"techdemo/factory/dialogue"
	"techdemo/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

type Dialogue struct {
}

func NewDialogue() *Dialogue {
	return &Dialogue{}
}

func (d *Dialogue) OnInteractEvent(w donburi.World, event events.Input) {

	if event != events.InputInteract {
		return
	}

	dialogueEntry, _ := tags.Dialogue.First(w)

	if dialogueEntry == nil {
		return
	}

	text, _ := transform.FindChildWithComponent(dialogueEntry, components.TextAnimation)
	animation := components.TextAnimation.Get(text)
	if animation.IsFinished() {
		dialogue.CloseDialogue(w)
	} else {
		animation.Skip()
	}
}
