package systems

import (
	"testing"

	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/events"

	"github.com/stretchr/testify/assert"
	"github.com/yohamta/donburi/features/math"
)

func TestToCardinalInput(t *testing.T) {

	cases := []struct {
		name     string
		v        math.Vec2
		expected events.Input
	}{
		{
			"cardinal left",
			constants.Left,
			events.InputMoveLeft,
		},
		{
			"cardinal right",
			constants.Right,
			events.InputMoveRight,
		},
		{
			"cardinal up",
			constants.Up,
			events.InputMoveUp,
		},
		{
			"cardinal down",
			constants.Down,
			events.InputMoveDown,
		},
		{
			"nne = up",
			math.NewVec2(1, -2),
			events.InputMoveUp,
		},
		{
			"ene = right",
			math.NewVec2(2, -1),
			events.InputMoveRight,
		},
		{
			"ese = right",
			math.NewVec2(2, 1),
			events.InputMoveRight,
		},
		{
			"sse = down",
			math.NewVec2(1, 2),
			events.InputMoveDown,
		},
		{
			"ssw = down",
			math.NewVec2(-1, 2),
			events.InputMoveDown,
		},
		{
			"wsw = left",
			math.NewVec2(-2, 1),
			events.InputMoveLeft,
		},
		{
			"wnw = left",
			math.NewVec2(-2, -1),
			events.InputMoveLeft,
		},

		{
			"nnw = up",
			math.NewVec2(-1, -2),
			events.InputMoveUp,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, toCardinalInput(tt.v))
		})
	}
}
