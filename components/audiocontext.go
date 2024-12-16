package components

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/yohamta/donburi"
)

var AudioContext = donburi.NewComponentType[audio.Context]()
