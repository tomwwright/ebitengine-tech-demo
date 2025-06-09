package archetypes

import (
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
)

type Archetype []donburi.IComponentType

func (a Archetype) Create(w donburi.World) *donburi.Entry {
	entity := w.Create(a...)
	return w.Entry(entity)
}

var (
	MusicPlayer     = Archetype{tags.MusicPlayer, components.AudioPlayer, components.Asset}
	Interaction     = Archetype{components.Transform, components.Object, components.Interaction}
	Assets          = Archetype{tags.Assets, components.Assets, components.AudioContext}
	Camera          = Archetype{tags.Camera, components.Transform, components.Movement, components.Camera}
	ScreenContainer = Archetype{tags.ScreenContainer, components.Transform}
	State           = Archetype{tags.State, components.State}
)
