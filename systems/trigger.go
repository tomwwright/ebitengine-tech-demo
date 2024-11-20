package systems

import (
	"techdemo/components"
	"techdemo/events"
	"techdemo/tags"

	"github.com/yohamta/donburi"
)

func OnMovementFinishedForTriggers(w donburi.World, entry *donburi.Entry) {
	if entry.HasComponent(tags.Player) {
		object := components.Object.Get(entry)

		if collision := object.Check(0, 0, tags.ResolvTagInteractive); collision != nil {
			other := collision.Objects[0]

			// trigger if we are overlapped in the same tile as the collider
			if object.Position == other.Position {
				events.TriggerEvent.Publish(w, other.Data.(*donburi.Entry))
			}
		}
	}
}
