package components

import (
	"github.com/yohamta/donburi"
)

type TextAnimationData struct {
	Text       string
	Speed      float32 // characters per second
	Characters float32 // number of characters currently displayed
}

var TextAnimation = donburi.NewComponentType[TextAnimationData]()

func (t *TextAnimationData) IsFinished() bool {
	return int(t.Characters) > len(t.Text)
}
