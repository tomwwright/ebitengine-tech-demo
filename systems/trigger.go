package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
)

func OnMovementFinishedForTriggers(w donburi.World, entry *donburi.Entry) {
	if entry.HasComponent(tags.Player) {
		object := components.Object.Get(entry)

		if collision := object.Check(0, 0, tags.ResolvTagInteractive); collision != nil {
			other := collision.Objects[0]
			position := object.Position.Sub(resolv.Vector(object.TransformOffset))

			// trigger if we are overlapped in the same tile as the collider
			if position == other.Position {
				events.TriggerEvent.Publish(w, other.Data.(*donburi.Entry))
			}
		}
	}
}
