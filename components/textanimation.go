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

func (t *TextAnimationData) Update(d float32) {
	t.Characters += t.Speed * d
}

func (t *TextAnimationData) Skip() {
	t.Characters = float32(len(t.Text))
}
