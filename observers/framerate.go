package observers

import (
	"fmt"
	"time"

	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/yohamta/donburi"
)

func SetupFrameRateObserver(w donburi.World) {
	counter := FrameRateCounter{}
	events.UpdateEvent.Subscribe(w, counter.OnUpdate)
}

type FrameRateCounter struct {
	frames      int
	accumulated int
	previous    time.Time
}

func (c FrameRateCounter) OnUpdate(w donburi.World, duration time.Duration) {
	c.accumulated += int(time.Now().Sub(c.previous).Milliseconds())
	c.frames++
	if c.frames > 240 {
		fmt.Printf("Average render time: %dms (%d frames)\n", c.accumulated/c.frames, c.frames)
		c.accumulated = 0
		c.frames = 0
	}
}
