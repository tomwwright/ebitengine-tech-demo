package interactions

import (
	"techdemo/events"
	"techdemo/interactions/yaml"
	"techdemo/sequences"

	"github.com/yohamta/donburi"
)

type Director struct {
	RunnableManager sequences.RunnableManager
	Interactions    *yaml.Interactions
}

func NewDirector() *Director {
	return &Director{
		RunnableManager: sequences.RunnableManager{},
	}
}

func (d *Director) SetInteractions(i *yaml.Interactions) {
	d.Interactions = i
}

func (d *Director) OnInteractionEvent(w donburi.World, event events.Interaction) {
	steps := d.Interactions.Interactions[event.Name]
	if steps == nil {
		return
	}
	sequence := constructSequence(w, steps)
	d.RunnableManager.Start(sequence)
}
