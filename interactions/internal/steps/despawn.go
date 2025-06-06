package steps

import (
	"fmt"

	"github.com/solarlune/resolv"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

type DespawnStep struct {
	World donburi.World
	Name  string
}

func (ds *DespawnStep) Run(done sequences.Done) {
	defer done()

	entry := findInteractionByName(ds.World, ds.Name)
	if entry != nil {
		findAndRemoveOverlappingObjectsByEntry(ds.World, entry)
	}

	fmt.Printf("DespawnStep: %s => %+v\n", ds.Name, entry)
}

func findInteractionByName(w donburi.World, name string) (entry *donburi.Entry) {
	components.Interaction.Each(w, func(e *donburi.Entry) {
		i := components.Interaction.Get(e)
		if i.Name == name {
			entry = e
		}
	})
	return entry
}

// findAndRemoveOverlappingObjectsByEntry locates Objects in the World that are in the same location
// and have the same size as the given entry
func findAndRemoveOverlappingObjectsByEntry(w donburi.World, entry *donburi.Entry) {
	object := components.Object.Get(entry)
	position := object.Position
	size := object.Size

	removals := []*resolv.Object{}
	components.Object.Each(w, func(e *donburi.Entry) {
		object := components.Object.Get(e)
		if object.Position == position && object.Size == size {
			removals = append(removals, object)
		}
	})

	for _, object := range removals {
		fmt.Printf("Removing %+v\n", object)
		object.Space.Remove(object)
		entry := object.Data.(*donburi.Entry)
		entity := entry.Entity()

		transform.RemoveRecursive(entry)
		w.Remove(entity)
	}
}
