package components

import (
	"image/color"

	"github.com/tanema/gween"
	"github.com/tomwwright/ebitengine-tech-demo/tween"
	"github.com/yohamta/donburi"
)

var Tween = donburi.NewComponentType[gween.Tween]()

var TweenColor = donburi.NewComponentType[tween.Tween[color.Color]]()
