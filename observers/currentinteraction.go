package observers

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
)

func SetupCurrentInteractionObserver(w donburi.World) {
	events.InteractionEvent.Subscribe(w, OnInteractionEvent)
	events.InteractionFinishedEvent.Subscribe(w, OnInteractionFinishedEvent)
}

func OnInteractionEvent(w donburi.World, event events.Interaction) {
	setCurrentInteraction(w, event.Target)
}

func OnInteractionFinishedEvent(w donburi.World, event events.Interaction) {
	setCurrentInteraction(w, nil)
}

func setCurrentInteraction(w donburi.World, target *donburi.Entry) {
	player := tags.Player.MustFirst(w)
	components.Target.SetValue(player, target)
	fmt.Printf("setCurrentInteraction %+v", target)
}
