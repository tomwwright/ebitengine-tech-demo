package interactions

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/interactions/yaml"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

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

func (d *Director) OnTriggerEvent(w donburi.World, entry *donburi.Entry) {
	name := components.Interaction.Get(entry).Name
	steps := d.Interactions.Triggers[name]
	if steps == nil {
		return
	}
	sequence := constructSequence(w, steps)
	d.RunnableManager.Start(sequence)
}
