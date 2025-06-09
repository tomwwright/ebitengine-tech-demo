package components

import (
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/yohamta/donburi"
)

type AssetData struct {
	Asset assets.Asset
}

var Asset = donburi.NewComponentType[AssetData]()
