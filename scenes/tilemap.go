package scenes

import (
	"fmt"
	"image/color"
	"techdemo/components"
	"techdemo/config"
	"techdemo/systems"
	"techdemo/tilemap"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type TilemapScene struct {
	ecs     *ecs.ECS
	Tilemap *tilemap.Tilemap
}

func NewTilemapScene(filename string) (*TilemapScene, error) {
	tilemap, err := tilemap.LoadTilemap(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load tilemap from %s: %w", filename, err)
	}

	scene := &TilemapScene{
		ecs:     ecs.NewECS(donburi.NewWorld()),
		Tilemap: tilemap,
	}
	scene.ecs.AddRenderer(ecs.LayerDefault, systems.NewRender().Draw)

	constructTileSprites(scene)

	return scene, nil
}

func (s *TilemapScene) Update() {
	s.ecs.Update()
}

func (s *TilemapScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
	s.ecs.Draw(screen)
}

func constructTileSprites(s *TilemapScene) {
	tilemap := s.Tilemap.Map
	w := s.ecs.World
	for _, l := range tilemap.Layers {
		for i, t := range l.Tiles {
			if !t.Nil {

				entity := w.Create(transform.Transform, components.Sprite)
				entry := w.Entry(entity)

				transform := transform.Transform.Get(entry)
				x := float64(l.OffsetX + (i%tilemap.Width)*tilemap.TileWidth)
				y := float64(l.OffsetY + (i/tilemap.Width)*tilemap.TileHeight)
				scale := float64(config.Scale)
				transform.LocalPosition = math.NewVec2(x*scale, y*scale)
				transform.LocalScale = math.NewVec2(scale, scale)

				sprite := components.Sprite.Get(entry)
				gid := t.ID + t.Tileset.FirstGID - 1
				sprite.Image = s.Tilemap.Tiles[gid]
			}

		}
	}
}
