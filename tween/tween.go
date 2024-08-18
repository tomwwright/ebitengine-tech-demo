package tween

import (
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
)

type SliceTween struct {
	Tweens []*gween.Tween
}

func NewSliceTween(from []float32, to []float32, duration float32, easing ease.TweenFunc) *SliceTween {
	tweens := make([]*gween.Tween, len(from))
	for i := range from {
		tweens[i] = gween.New(from[i], to[i], duration, easing)
	}
	return &SliceTween{
		Tweens: tweens,
	}
}

func (s *SliceTween) Update(dt float32) (current []float32, isFinished bool) {
	current = make([]float32, len(s.Tweens))
	for i, tween := range s.Tweens {
		current[i], isFinished = tween.Update(dt)
	}
	return current, isFinished
}
