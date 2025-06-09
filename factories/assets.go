package factories

import (
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/yohamta/donburi"
)

func CreateAssets(w donburi.World, files fs.ReadFileFS) *donburi.Entry {
	e := archetypes.Assets.Create(w)
	components.AudioContext.Set(e, audio.NewContext(constants.AudioSampleRate))
	components.Assets.Set(e, &components.AssetsData{
		Assets: assets.NewFileSystemAssets(files),
	})
	return e
}
