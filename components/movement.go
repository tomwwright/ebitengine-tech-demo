package components

import (
	"techdemo/tween"

	"github.com/yohamta/donburi"
)

type MovementData struct {
	Tween *tween.Vec2Tween
}

var Movement = donburi.NewComponentType[MovementData]()
