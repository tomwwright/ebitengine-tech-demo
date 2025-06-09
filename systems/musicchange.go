package systems

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
)

var musicChangeQuery = donburi.NewQuery(filter.Contains(tags.MusicPlayer, components.AudioPlayer, components.Tween))

func UpdateMusicChange(ecs *ecs.ECS) {
	musicChangeQuery.Each(ecs.World, func(e *donburi.Entry) {
		t := components.Tween.Get(e)
		audio := components.AudioPlayer.Get(e)
		volume, isFinished := t.Update(constants.DeltaTime)
		audio.SetVolume(float64(volume))

		if isFinished {
			if volume > 0 {
				e.RemoveComponent(components.Tween)
			} else {
				e.Remove() // if we were tweening to zero, delete the music player
				audio.Close()
			}
		}
	})
}
