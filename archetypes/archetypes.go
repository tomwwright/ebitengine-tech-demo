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
	CameraContainer = Archetype{tags.CameraContainer, components.Transform}
	State           = Archetype{tags.State, components.State}
	Dialogue        = Archetype{tags.Dialogue, components.Transform, components.Sprite, components.AudioPlayer}
	Text            = Archetype{components.Transform, components.Text, components.TextAnimation}
	Player          = Archetype{tags.Player, components.Transform, components.Sprite, components.Object, components.Movement, components.Animation, components.CharacterAnimations, components.Target}
	Sprite          = Archetype{components.Transform, components.Sprite}
)
