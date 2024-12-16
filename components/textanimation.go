package components

import (
	"github.com/yohamta/donburi"
)

type TextAnimationData struct {
	Text       string
	Speed      float32     // characters per second
	Characters float32     // number of characters currently displayed
	OnTick     func(n int) // called when characters changes
}

var TextAnimation = donburi.NewComponentType[TextAnimationData]()

func (t *TextAnimationData) IsFinished() bool {
	return t.CharactersInt() > len(t.Text)
}

func (t *TextAnimationData) Update(d float32) {
	n := t.CharactersInt()
	t.Characters += t.Speed * d
	if t.CharactersInt() > n && t.OnTick != nil {
		t.OnTick(t.CharactersInt())
	}
}

func (t *TextAnimationData) Skip() {
	t.Characters = t.maxCharacters()
	if t.OnTick != nil {
		t.OnTick(t.CharactersInt())
	}
}

func (t *TextAnimationData) CharactersInt() int {
	return int(t.Characters)
}

func (t *TextAnimationData) maxCharacters() float32 {
	return float32(len(t.Text))
}
