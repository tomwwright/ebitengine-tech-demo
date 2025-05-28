package steps

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

	"github.com/yohamta/donburi"
)

type DebugStep struct {
	Text  string
	World donburi.World
}

func (ds *DebugStep) Run(done sequences.Done) {
	defer done()

	fmt.Printf("DebugStep: %s (State: %+v)\n", ds.Text, components.State.Get(components.State.MustFirst(ds.World)))

}
