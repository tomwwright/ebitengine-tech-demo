package systems

import (
	"fmt"
	"techdemo/components"
	"techdemo/constants"
	"techdemo/events"
	"techdemo/tags"

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
				events.DialogueEvent.Publish(w, events.Dialogue{
					Text: fmt.Sprintf("Interaction %+v\n", interaction),
				})
				fmt.Printf("Interaction %+v\n", interaction)
			}
		}

	}
}
