package interactions

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/interactions/internal/steps"
	"github.com/tomwwright/ebitengine-tech-demo/interactions/yaml"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

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
		} else if s.Teleport != nil {
			step = &steps.TeleportStep{
				To:    s.Teleport.To,
				World: w,
			}
		} else if s.Filter != nil {
			step = &steps.FilterStep{
				Name:  s.Filter.Name,
				World: w,
			}
		} else if s.Music != nil {
			step = &steps.MusicStep{
				Name:  s.Music.Name,
				World: w,
			}
		} else if s.Camera != nil {
			step = &steps.ZoomStep{
				Zoom:  s.Camera.Zoom,
				World: w,
			}
		} else {
			fmt.Printf("Unknown step: %+v\n", s)
		}

		sequence.Steps = append(sequence.Steps, step)
	}
	return sequence
}
