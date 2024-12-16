package components

import (
	"techdemo/assets"

	"github.com/yohamta/donburi"
)

type AssetProvider interface {
	GetAsset(asset assets.Asset) ([]byte, error)
}

type AssetsData struct {
	Assets AssetProvider
}

var Assets = donburi.NewComponentType[AssetsData]()
