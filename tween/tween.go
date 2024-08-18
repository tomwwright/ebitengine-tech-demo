package tween

import (
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi/features/math"
)

type Vec2Tween struct {
	Tweens [2]*gween.Tween
}

func NewVec2Tween(from math.Vec2, to math.Vec2, duration float32, easing ease.TweenFunc) *Vec2Tween {
	var tweens [2]*gween.Tween
	tweens[0] = gween.New(float32(from.X), float32(to.X), duration, easing)
	tweens[1] = gween.New(float32(from.Y), float32(to.Y), duration, easing)
	return &Vec2Tween{
		Tweens: tweens,
	}
}

func (s *Vec2Tween) Update(dt float32) (current math.Vec2, isFinished bool) {
	x, _ := s.Tweens[0].Update(dt)
	y, isFinished := s.Tweens[1].Update(dt)

	return math.NewVec2(float64(x), float64(y)), isFinished
}
