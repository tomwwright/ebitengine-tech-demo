package factories

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
)

func CreateMusicPlayer(w donburi.World, asset assets.Asset) (*donburi.Entry, error) {
	e := tags.Assets.MustFirst(w)
	context := components.AudioContext.Get(e)
	assets := components.Assets.Get(e)

	e = archetypes.MusicPlayer.Create(w)
	b, _ := assets.Assets.GetAsset(asset)
	stream, _ := vorbis.DecodeF32(bytes.NewReader(b))
	audio, err := context.NewPlayerF32(stream)
	if err != nil {
		return nil, fmt.Errorf("unable to create audio player: %w", err)
	}
	components.AudioPlayer.Set(e, audio)
	components.Asset.SetValue(e, components.AssetData{
		Asset: asset,
	})

	return e, nil
}
