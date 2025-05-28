package components

import (
	"github.com/tomwwright/ebitengine-tech-demo/tilemap"

	"github.com/yohamta/donburi"
)

type CharacterAnimationsData struct {
	Animations map[string]tilemap.Animation
}

var CharacterAnimations = donburi.NewComponentType[CharacterAnimationsData]()

func (c *CharacterAnimationsData) Add(animation tilemap.Animation) {
	if c.Animations == nil {
		c.Animations = map[string]tilemap.Animation{}
	}
	c.Animations[animation.Name] = animation
}
