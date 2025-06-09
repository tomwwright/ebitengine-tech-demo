package steps

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/tanema/gween"
	"github.com/tanema/gween/ease"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/components"
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

	q := donburi.NewQuery(filter.Contains(tags.MusicPlayer, components.AudioPlayer, components.Asset))

	if q.Count(s.World) > 1 {
		return
	}

	// TODO this needs to live elsewhere
	tracks := map[string]assets.Asset{
		"town":   assets.AssetAudioMusic,
		"forest": assets.AssetAudioMusicForest,
	}
	asset := tracks[s.Name]

	// check current track isn't already playing
	e := tags.MusicPlayer.MustFirst(s.World)
	if components.Asset.Get(e).Asset == asset {
		return
	}

	// tween current audio player to zero
	audio := components.AudioPlayer.Get(e)
	volume := audio.Volume()
	e.AddComponent(components.Tween)
	components.Tween.Set(e, gween.New(float32(volume), 0, 1.0, ease.Linear))

	e = tags.Assets.MustFirst(s.World)

	context := components.AudioContext.Get(e)
	assets := components.Assets.Get(e)

	// create new music player and tween up to current volume
	e = s.World.Entry(s.World.Create(tags.MusicPlayer, components.AudioPlayer, components.Asset))

	b, _ := assets.Assets.GetAsset(asset)
	stream, _ := vorbis.DecodeF32(bytes.NewReader(b))
	audio, _ = context.NewPlayerF32(stream)
	audio.SetVolume(0)
	audio.Play()
	components.AudioPlayer.Set(e, audio)
	e.AddComponent(components.Tween)
	components.Tween.Set(e, gween.New(0, float32(volume), 1.0, ease.Linear))
	components.Asset.SetValue(e, components.AssetData{
		Asset: asset,
	})

	fmt.Printf("Music: %s\n", s.Name)
}
