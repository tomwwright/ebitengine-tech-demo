package steps

import (
	"fmt"
	"techdemo/sequences"
)

type DebugStep struct {
	Text string
}

func (ds *DebugStep) Run(done sequences.Done) {
	defer done()

	fmt.Printf("DebugStep: %s\n", ds.Text)
}
