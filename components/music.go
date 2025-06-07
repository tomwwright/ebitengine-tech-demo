package components

import (
	"bytes"

	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
)

type MusicData struct {
	Volume float64
	tracks map[string]assets.Asset
	self   *donburi.Entry
}

var Music = donburi.NewComponentType[MusicData]()

func NewMusic(entry *donburi.Entry) *MusicData {
	return &MusicData{
		Volume: 0.3,
		tracks: map[string]assets.Asset{
			"town":   assets.AssetAudioMusic,
			"forest": assets.AssetAudioMusicForest,
		},
		self: entry,
	}
}

func (m *MusicData) ChangeTrack(track string) bool {
	e := tags.Assets.MustFirst(m.self.World)

	context := AudioContext.Get(e)
	assets := Assets.Get(e)

	asset := m.tracks[track]
	if asset == "" {
		return false
	}

	audioPlayer := AudioPlayer.Get(m.self)
	if audioPlayer != nil {
		audioPlayer.Close()
	}

	b, _ := assets.Assets.GetAsset(asset)
	stream, _ := vorbis.DecodeF32(bytes.NewReader(b))
	audioPlayer, _ = context.NewPlayerF32(stream)
	audioPlayer.SetVolume(m.Volume)
	AudioPlayer.Set(m.self, audioPlayer)

	audioPlayer.Play()

	return true
}
