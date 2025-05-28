package steps

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/interactions/yaml"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

	"github.com/yohamta/donburi"
)

type StateStep struct {
	yaml.State
	World donburi.World
}

func (s *StateStep) Run(done sequences.Done) {
	defer done()

	e := components.State.MustFirst(s.World)
	state := components.State.Get(e)

	switch s.Action {
	case "increment":
		state.Increment(s.Key)
	case "true":
		state.SetTrue(s.Key)
	default:
		state.Set(s.Key, s.Value)
	}
}
