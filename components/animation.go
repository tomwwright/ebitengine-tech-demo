package components

import (
	"techdemo/tilemap"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
)

type Status int

const (
	Playing = iota
	Paused
)

type AnimationData struct {
	Name      string
	Frames    []*ebiten.Image
	Durations []time.Duration
	index     int
	timer     time.Duration
	status    Status
}

func NewAnimation(frames []*ebiten.Image, durations []time.Duration) *AnimationData {
	return &AnimationData{
		Frames:    frames,
		Durations: durations,
		index:     0,
		timer:     0,
		status:    Playing,
	}
}

func NewAnimationFromTilemapAnimation(animation tilemap.Animation) *AnimationData {
	durations := []time.Duration{}
	frames := []*ebiten.Image{}
	for _, frame := range animation.Frames {
		durations = append(durations, frame.Duration)
		frames = append(frames, frame.Image)
	}

	return &AnimationData{
		Name:      animation.Name,
		Durations: durations,
		Frames:    frames,
		index:     0,
		timer:     0,
		status:    Playing,
	}
}

func (anim *AnimationData) Clone() *AnimationData {
	new := *anim
	return &new
}

func (anim *AnimationData) IsEnd() bool {
	if anim.status == Paused && anim.index == anim.Length()-1 {
		return true
	}
	return false
}

func (anim *AnimationData) Update(elapsedTime time.Duration) {
	if anim.status != Playing || anim.Length() <= 1 {
		return
	}

	anim.timer += elapsedTime
	currentFrameDuration := anim.FrameDuration()
	if anim.timer >= currentFrameDuration {
		anim.timer -= currentFrameDuration
		anim.index++
	}

	if anim.index > len(anim.Frames)-1 {
		anim.index = 0
	}
}

func (anim *AnimationData) Image() *ebiten.Image {
	if anim.IsNil() {
		return nil
	}
	return anim.Frames[anim.index]
}

func (anim *AnimationData) Status() Status {
	return anim.status
}

func (anim *AnimationData) Pause() {
	anim.status = Paused
}

func (anim *AnimationData) Index() int {
	return anim.index
}

func (anim *AnimationData) Length() int {
	return len(anim.Frames)
}

func (anim *AnimationData) IsNil() bool {
	return anim.Length() == 0
}

func (anim *AnimationData) FrameDuration() time.Duration {
	if anim.IsNil() {
		return 0
	}
	return anim.Durations[anim.index]
}

func (anim *AnimationData) GoToFrame(index int) {
	anim.index = index
	anim.timer = 0
}

func (anim *AnimationData) PauseAtEnd() {
	anim.index = anim.Length() - 1
	anim.timer = anim.FrameDuration()
	anim.Pause()
}

func (anim *AnimationData) PauseAtStart() {
	anim.index = 0
	anim.timer = 0
	anim.Pause()
}

func (anim *AnimationData) Resume() {
	anim.status = Playing
}

var Animation = donburi.NewComponentType[AnimationData]()
