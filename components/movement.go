package components

import (
	"techdemo/tween"

	"github.com/yohamta/donburi"
)

type MovementData struct {
	Tween *tween.SliceTween
}

var Movement = donburi.NewComponentType[MovementData]()
