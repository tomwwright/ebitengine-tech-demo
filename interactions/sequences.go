package interactions

import (
	"fmt"
	"techdemo/interactions/internal/steps"
	"techdemo/interactions/yaml"
	"techdemo/sequences"

	"github.com/yohamta/donburi"
)

func constructSequence(w donburi.World, interactions *yaml.Interactions, name string) *sequences.Sequence {
	sequence := &sequences.Sequence{
		Steps: []sequences.Runnable{},
	}

	interaction := interactions.Interactions[name]
	if interaction == nil {
		return sequence
	}

	for _, s := range interaction {
		var step sequences.Runnable
		if s.Debug != nil {
			step = &steps.DebugStep{
				Text:  fmt.Sprintf("%s: %s", name, s.Debug.Text),
				World: w,
			}
		} else if s.Dialogue != nil {
			step = &steps.DialogueStep{
				Text:  s.Dialogue.Text,
				World: w,
			}
		} else if s.Despawn != nil {
			step = &steps.DespawnStep{
				Name:  s.Despawn.Name,
				World: w,
			}
		} else if s.State != nil {
			step = &steps.StateStep{
				State: *s.State,
				World: w,
			}
		} else {
			fmt.Printf("Unknown step in %s: %+v\n", name, s)
		}

		sequence.Steps = append(sequence.Steps, step)
	}
	return sequence
}
