package systems

import (
	"github.com/solarlune/resolv"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/yohamta/donburi"
)

func OnMovementFinishedForTriggers(w donburi.World, entry *donburi.Entry) {
	if entry.HasComponent(tags.Player) {
		object := components.Object.Get(entry)

		if collision := object.Check(0, 0, tags.ResolvTagInteractive); collision != nil {
			other := collision.Objects[0]
			if isContainedBy(object, other) {
				events.TriggerEvent.Publish(w, other.Data.(*donburi.Entry))
			}
		}
	}
}

func isContainedBy(object *resolv.Object, other *resolv.Object) bool {
	return object.Position.X >= other.Position.X &&
		object.Position.X+object.Size.X <= other.Position.X+other.Size.X &&
		object.Position.Y >= other.Position.Y &&
		object.Position.Y+object.Size.Y <= other.Position.Y+other.Size.Y
}
