package constants

import (
	"time"

	"github.com/yohamta/donburi/features/math"
)

const (
	ScreenWidth       = 512
	ScreenHeight      = 786
	TileSize          = 16
	Scale             = 4
	DeltaTime         = 1.0 / 60.0
	DeltaTimeDuration = time.Second / 60.0
	LayerUI           = 8
	AudioSampleRate   = 22000
	LineSpacing       = 16
)

var Up = math.NewVec2(0, -1)
var Down = math.NewVec2(0, 1)
var Left = math.NewVec2(-1, 0)
var Right = math.NewVec2(1, 0)
var Zero = math.NewVec2(0, 0)
