package steps

import (
	"fmt"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/yohamta/donburi"
)

type ZoomStep struct {
	World donburi.World
	Zoom  float64
}

func (s *ZoomStep) Run(done sequences.Done) {
	defer done()
	w := s.World

	_, exists := tags.ZoomChange.First(w)
	if exists {
		return
	}

	c := components.Camera.MustFirst(w)
	transform := components.Transform.Get(c)

	if transform.LocalScale.X == s.Zoom {
		return
	}

	entity := s.World.Create(tags.ZoomChange, components.Target, components.Tween)
	e := s.World.Entry(entity)

	components.Target.SetValue(e, c)
	components.Tween.Set(e, gween.New(float32(transform.LocalScale.X), float32(s.Zoom), 1.0, ease.InQuad))

	fmt.Printf("ZoomStep: %d\n", s.Zoom)
}
