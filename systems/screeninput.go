package systems

import (
	"math"

	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/events"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	dmath "github.com/yohamta/donburi/features/math"
)

func toCardinalInput(v dmath.Vec2) events.Input {
	d := v.Normalized()
	if math.Abs(d.Y) > math.Abs(d.X) {
		if d.Y > 0 {
			return events.InputMoveDown
		} else {
			return events.InputMoveUp
		}
	} else {
		if d.X > 0 {
			return events.InputMoveRight
		} else {
			return events.InputMoveLeft
		}
	}
}

type StrokeSource interface {
	IsJustReleased() bool
	Position() (int, int)
}

type TouchStrokeSource struct {
	touchId ebiten.TouchID
}

func (t *TouchStrokeSource) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.touchId)
}

func (t *TouchStrokeSource) Position() (int, int) {
	return ebiten.TouchPosition(t.touchId)
}

type MouseStrokeSource struct {
	button ebiten.MouseButton
}

func (m *MouseStrokeSource) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(m.button)
}

func (m *MouseStrokeSource) Position() (int, int) {
	return ebiten.CursorPosition()
}

type Stroke struct {
	source    StrokeSource
	startedAt dmath.Vec2
	position  dmath.Vec2
	hasMoved  bool
}

func NewStroke(source StrokeSource) *Stroke {
	x, y := source.Position()
	return &Stroke{
		source:    source,
		startedAt: dmath.NewVec2(float64(x), float64(y)),
		hasMoved:  false,
	}
}

func (s *Stroke) GetCurrentInput() events.Input {
	s.update()
	m := s.Movement()
	if m.Magnitude() < constants.TileSize {
		return events.InputMoveNone
	} else {
		s.hasMoved = true // record that this touch has moved away from initial touch
		return toCardinalInput(m)
	}
}

func (s *Stroke) HasMoved() bool {
	return s.hasMoved
}

func (s *Stroke) Movement() dmath.Vec2 {
	return s.position.Sub(s.startedAt)
}

func (s *Stroke) update() {
	x, y := s.source.Position()
	s.position.X = float64(x)
	s.position.Y = float64(y)
}

type ScreenInput struct {
	stroke    *Stroke
	buffer    []ebiten.TouchID
	lastEvent events.Input
}

func NewScreenInput() *ScreenInput {
	return &ScreenInput{
		buffer:    []ebiten.TouchID{},
		lastEvent: events.InputMoveNone,
	}
}

func (input *ScreenInput) publishEvent(w donburi.World, event events.Input) {
	if event != input.lastEvent {
		events.InputEvent.Publish(w, event)
		input.lastEvent = event
	}
}

func (input *ScreenInput) CheckForPress() {
	input.buffer = inpututil.AppendJustPressedTouchIDs(input.buffer[:0])
	if len(input.buffer) != 0 {
		input.stroke = NewStroke(&TouchStrokeSource{touchId: input.buffer[0]})
		return
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		input.stroke = NewStroke(&MouseStrokeSource{button: ebiten.MouseButton0})
	}
}

func (input *ScreenInput) OnStrokeReleased(w donburi.World) {
	input.publishEvent(w, events.InputMoveNone)
	if !input.stroke.HasMoved() {
		input.publishEvent(w, events.InputInteract)
	}
	input.stroke = nil
}

func (input *ScreenInput) SendCurrentInput(w donburi.World) {
	i := input.stroke.GetCurrentInput()
	input.publishEvent(w, i)
}

func (input *ScreenInput) Update(ecs *ecs.ECS) {
	w := ecs.World

	if input.stroke == nil {
		input.CheckForPress()
	} else if input.stroke.source.IsJustReleased() {
		input.OnStrokeReleased(w)
	} else {
		input.SendCurrentInput(w)
	}
}
