package components

import (
	"github.com/tomwwright/ebitengine-tech-demo/assets"

	"github.com/yohamta/donburi"
)

type AssetProvider interface {
	GetAsset(asset assets.Asset) ([]byte, error)
}

type AssetsData struct {
	Assets AssetProvider
}

var Assets = donburi.NewComponentType[AssetsData]()
