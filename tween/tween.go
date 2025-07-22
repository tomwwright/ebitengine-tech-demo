package tween

import (
	"fmt"
	"image/color"
	"time"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/yohamta/donburi/features/math"
)

type Vec2Tween struct {
	Tweens   [2]*gween.Tween
	From     math.Vec2
	To       math.Vec2
	OnFinish func()
}

func NewVec2Tween(from math.Vec2, to math.Vec2, duration float32, easing ease.TweenFunc) *Vec2Tween {
	var tweens [2]*gween.Tween
	tweens[0] = gween.New(float32(from.X), float32(to.X), duration, easing)
	tweens[1] = gween.New(float32(from.Y), float32(to.Y), duration, easing)
	return &Vec2Tween{
		Tweens: tweens,
		From:   from,
		To:     to,
	}
}

func (s *Vec2Tween) Update(dt float32) (current math.Vec2, isFinished bool) {
	x, _ := s.Tweens[0].Update(dt)
	y, isFinished := s.Tweens[1].Update(dt)

	if isFinished && s.OnFinish != nil {
		s.OnFinish()
	}

	return math.NewVec2(float64(x), float64(y)), isFinished
}

type Tween[T any] struct {
	gween.Tween
	From     T
	To       T
	OnFinish func()
}

func NewTween[T any](from T, to T, duration time.Duration, easing ease.TweenFunc) Tween[T] {
	return Tween[T]{
		Tween: *gween.New(0.0, 1.0, float32(duration.Seconds()), easing),
		From:  from,
		To:    to,
	}
}

func (t *Tween[T]) Update(dt float32) (current T, isFinished bool) {
	d, isFinished := t.Tween.Update(dt)

	from := any(t.From)
	to := any(t.To)

	switch from.(type) {
	case color.Color:
		current = TweenColors(from.(color.Color), to.(color.Color), d).(T)
	default:
		panic(fmt.Errorf("unable to tween type: %+v -> %+v", from, to))
	}

	if isFinished && t.OnFinish != nil {
		t.OnFinish()
	}

	return current, isFinished
}

func TweenColors(from color.Color, to color.Color, amount float32) color.Color {
	fromR, fromG, fromB, fromA := from.RGBA()
	toR, toG, toB, toA := to.RGBA()

	r := int(fromR) + int((float32(uint8(toR))-float32(uint8(fromR)))*amount)
	g := int(fromG) + int((float32(uint8(toG))-float32(uint8(fromG)))*amount)
	b := int(fromB) + int((float32(uint8(toB))-float32(uint8(fromB)))*amount)
	a := int(fromA) + int((float32(uint8(toA))-float32(uint8(fromA)))*amount)

	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}
