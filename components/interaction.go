package components

import (
	"github.com/yohamta/donburi"
)

type InteractionData struct {
	Name string
}

var Interaction = donburi.NewComponentType[InteractionData]()
