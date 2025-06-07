package steps

import (
	"fmt"
	"image/color"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"

	"github.com/yohamta/donburi"
)

type FilterStep struct {
	World donburi.World
	Name  string
}

var filters = map[string]color.Color{
	"day":   color.White,
	"night": color.RGBA{150, 160, 220, 255},
}

func (s *FilterStep) Run(done sequences.Done) {
	defer done()

	color := filters[s.Name]

	e := components.Camera.MustFirst(s.World)
	camera := components.Camera.Get(e)
	camera.Color = color

	fmt.Printf("FilterStep: %s => %+v\n", s.Name, color)
}
