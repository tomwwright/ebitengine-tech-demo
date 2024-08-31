package components

import (
	"techdemo/tween"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type MovementData struct {
	Tween         *tween.Vec2Tween
	LastDirection math.Vec2
}

var Movement = donburi.NewComponentType[MovementData]()
