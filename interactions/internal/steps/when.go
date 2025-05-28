package steps

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/interactions/yaml"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

	"github.com/yohamta/donburi"
)

type WhenStep struct {
	Conditions []yaml.Condition
	Steps      sequences.Runnable
	Else       sequences.Runnable
	World      donburi.World
}

func (w *WhenStep) Run(done sequences.Done) {
	if evaluateAnd(w.World, w.Conditions) {
		w.Steps.Run(done)
	} else {
		w.Else.Run(done)
	}
}

func evaluateAnd(w donburi.World, conditions []yaml.Condition) bool {
	for _, c := range conditions {
		if evaluateCondition(w, c) == false {
			return false
		}
	}
	return true
}

func evaluateOr(w donburi.World, conditions []yaml.Condition) bool {
	isTrue := false
	for _, c := range conditions {
		if evaluateCondition(w, c) {
			isTrue = true
		}
	}
	return isTrue
}

func evaluateCondition(w donburi.World, condition yaml.Condition) bool {
	if condition.Or != nil {
		return evaluateOr(w, condition.Or)
	} else if condition.State != nil {
		return evaluateStateCondition(w, *condition.State)
	} else {
		return false
	}
}

func evaluateStateCondition(w donburi.World, condition yaml.StateCondition) bool {
	e := components.State.MustFirst(w)
	state := components.State.Get(e)
	return state.Get(condition.Key) == condition.Value
}
