package components

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/yohamta/donburi"
)

var AudioPlayer = donburi.NewComponentType[audio.Player]()
