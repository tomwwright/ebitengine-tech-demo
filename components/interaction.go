package components

import (
	"github.com/yohamta/donburi"
)

type InteractionData struct {
	Payload string
}

var Interaction = donburi.NewComponentType[InteractionData]()
