package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/factories/dialogue"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

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
