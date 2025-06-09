package steps

import (
	"fmt"

	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/factories"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type MusicStep struct {
	World donburi.World
	Name  string
}

func (s *MusicStep) Run(done sequences.Done) {
	defer done()
	w := s.World
	q := donburi.NewQuery(filter.Contains(archetypes.MusicPlayer...))

	if q.Count(w) > 1 {
		return
	}

	// TODO this needs to live elsewhere
	tracks := map[string]assets.Asset{
		"town":   assets.AssetAudioMusic,
		"forest": assets.AssetAudioMusicForest,
	}
	asset := tracks[s.Name]

	// check current track isn't already playing
	e := tags.MusicPlayer.MustFirst(w)
	if components.Asset.Get(e).Asset == asset {
		return
	}

	// tween current audio player to zero
	audio := components.AudioPlayer.Get(e)
	volume := audio.Volume()
	e.AddComponent(components.Tween)
	components.Tween.Set(e, gween.New(float32(volume), 0, 1.0, ease.Linear))

	// create new music player and tween up to current volume
	e, err := factories.CreateMusicPlayer(w, asset)
	if err != nil {
		panic(err)
	}
	audio = components.AudioPlayer.Get(e)
	audio.SetVolume(0)
	audio.Play()
	e.AddComponent(components.Tween)
	components.Tween.Set(e, gween.New(0, float32(volume), 1.0, ease.Linear))

	fmt.Printf("Music: %s\n", s.Name)
}
