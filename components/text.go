package components

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
)

type TextData struct {
	Text  string
	Font  text.Face
	Layer int
}

var Text = donburi.NewComponentType[TextData]()
