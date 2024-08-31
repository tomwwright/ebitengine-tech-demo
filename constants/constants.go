package constants

import (
	"time"

	"github.com/yohamta/donburi/features/math"
)

const (
	ScreenWidth       = 640
	ScreenHeight      = 480
	TileSize          = 16
	Scale             = 2
	DeltaTime         = 1.0 / 60.0
	DeltaTimeDuration = time.Second / 60.0
)

var Up = math.NewVec2(0, -1)
var Down = math.NewVec2(0, 1)
var Left = math.NewVec2(-1, 0)
var Right = math.NewVec2(1, 0)
var Zero = math.NewVec2(0, 0)