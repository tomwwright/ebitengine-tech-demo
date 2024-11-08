package interactions

import (
	"fmt"
	"techdemo/interactions/internal/steps"
	"techdemo/interactions/yaml"
	"techdemo/sequences"

	"github.com/yohamta/donburi"
)

func constructSequence(w donburi.World, stepsList []yaml.Step) *sequences.Sequence {
	sequence := &sequences.Sequence{
		Steps: []sequences.Runnable{},
	}

	for _, s := range stepsList {
		var step sequences.Runnable
		if s.Debug != nil {
			step = &steps.DebugStep{
				Text:  s.Debug.Text,
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
		} else if s.When != nil {
			step = &steps.WhenStep{
				Conditions: s.When.Conditions,
				Steps:      constructSequence(w, s.When.Steps),
				Else:       constructSequence(w, s.When.Else),
				World:      w,
			}
		} else {
			fmt.Printf("Unknown step: %+v\n", s)
		}

		sequence.Steps = append(sequence.Steps, step)
	}
	return sequence
}
