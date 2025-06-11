package constants

import (
	"image/color"
	"time"

	"github.com/yohamta/donburi/features/math"
)

const (
	ScreenWidth       = 512
	ScreenHeight      = 786
	TileSize          = 16
	Scale             = 4
	Width             = ScreenWidth / Scale
	Height            = ScreenHeight / Scale
	DeltaTime         = 1.0 / 60.0
	DeltaTimeDuration = time.Second / 60.0
	LayerUI           = 99
	AudioSampleRate   = 22000
	LineSpacing       = 16
)

var Up = math.NewVec2(0, -1)
var Down = math.NewVec2(0, 1)
var Left = math.NewVec2(-1, 0)
var Right = math.NewVec2(1, 0)
var Zero = math.NewVec2(0, 0)

var White = color.RGBA{255, 255, 255, 255}
var Blueish = color.RGBA{150, 160, 220, 255}
