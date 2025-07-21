package systems

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

func OnInteractEvent(w donburi.World, input events.Input) {
	if input == events.InputInteract {
		q := donburi.NewQuery(filter.Contains(tags.Player, components.Movement, components.Object))
		player, ok := q.First(w)
		if !ok {
			return
		}

		object := components.Object.Get(player)
		movement := components.Movement.Get(player)

		if movement.LastDirection == constants.Zero {
			return
		}

		d := float64(constants.TileSize / 2)
		dv := movement.LastDirection.MulScalar(d)

		if collision := object.Check(dv.X, dv.Y, tags.ResolvTagInteractive); collision != nil {
			fmt.Printf("%+v %+v\n", dv, collision)
			entry := components.ResolveObjectEntry(collision.Objects[0])

			if entry.HasComponent(components.Interaction) {
				interaction := components.Interaction.Get(entry)
				events.InteractionEvent.Publish(w, events.Interaction{
					Name:   interaction.Name,
					Target: entry,
				})
				fmt.Printf("Interaction %+v\n", interaction)
			}
		}

	}
}
