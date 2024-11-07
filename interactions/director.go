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
	sequence := constructSequence(w, d.Interactions, event.Name)
	d.RunnableManager.Start(sequence)
}
