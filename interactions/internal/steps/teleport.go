package steps

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

type TeleportStep struct {
	World donburi.World
	To    string
}

func (ts *TeleportStep) Run(done sequences.Done) {
	defer done()

	entry := findInteractionByName(ts.World, ts.To)
	if entry != nil {
		teleportPlayerToEntry(ts.World, entry)
	}

	fmt.Printf("TeleportStep: %s => %+v\n", ts.To, entry)
}

func teleportPlayerToEntry(w donburi.World, entry *donburi.Entry) {
	player := tags.Player.MustFirst(w)

	t := transform.GetTransform(player)
	destination := transform.GetTransform(entry)

	t.LocalPosition = destination.LocalPosition

	// stop any current movement
	m := components.Movement.Get(player)
	m.Stop()
}
