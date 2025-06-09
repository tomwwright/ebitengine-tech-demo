package steps

import (
	"fmt"
	"image/color"
	"time"

	"github.com/tanema/gween/ease"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/tomwwright/ebitengine-tech-demo/tween"

	"github.com/yohamta/donburi"
)

type FilterStep struct {
	World donburi.World
	Name  string
}

var filters = map[string]color.Color{
	"day":   constants.White,
	"night": constants.Blueish,
}

func (s *FilterStep) Run(done sequences.Done) {
	defer done()

	_, exists := tags.FilterChange.First(s.World)
	if exists {
		return
	}

	c := components.Camera.MustFirst(s.World)
	camera := components.Camera.Get(c)
	color := filters[s.Name]

	if camera.Color == color {
		return
	}

	entity := s.World.Create(tags.FilterChange, components.Target, components.TweenColor)
	e := s.World.Entry(entity)

	components.Target.SetValue(e, c)
	components.TweenColor.SetValue(e, tween.NewTween(camera.Color, color, time.Second, ease.InQuad))

	fmt.Printf("FilterStep: %s => %+v\n", s.Name, color)
}
