package steps

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/sequences"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/yohamta/donburi"
)

type MusicStep struct {
	World donburi.World
	Name  string
}

func (s *MusicStep) Run(done sequences.Done) {
	defer done()

	e := tags.MusicPlayer.MustFirst(s.World)
	music := components.Music.Get(e)
	music.ChangeTrack(s.Name)

	fmt.Printf("Music: %s\n", s.Name)
}
