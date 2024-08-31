package tilemap

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/samber/lo"
)

type Tilemap struct {
	Map        *tiled.Map
	Filename   string
	Tiles      map[uint32]*ebiten.Image
	Tilesets   map[string]*ebiten.Image
	Animations []Animation
}

func LoadTilemap(filename string) (*Tilemap, error) {
	m := &Tilemap{
		Filename: filename,
	}
	err := m.loadMap()
	if err != nil {
		return nil, err
	}

	err = m.loadTilesets()
	if err != nil {
		return nil, err
	}

	m.loadAnimations()

	return m, nil
}

func (m *Tilemap) loadMap() error {
	tiledMap, err := tiled.LoadFile(m.Filename)
	if err != nil {
		return fmt.Errorf("error parsing map: %w", err)
	}
	m.Map = tiledMap
	return nil
}

func (m *Tilemap) loadTilesets() error {
	m.Tiles = map[uint32]*ebiten.Image{}
	m.Tilesets = map[string]*ebiten.Image{}
	dir := filepath.Dir(m.Filename)
	for _, tileset := range m.Map.Tilesets {
		filename := filepath.Join(dir, tileset.Image.Source)
		fmt.Println("loading tileset", tileset.Name, "from", filename)
		source, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("error loading tileset image file: %w", err)
		}

		img, _, err := image.Decode(bytes.NewReader(source))
		if err != nil {
			return fmt.Errorf("error decoding tileset image file: %w", err)
		}
		tilesetImage := ebiten.NewImageFromImage(img)
		m.Tilesets[tileset.Name] = tilesetImage

		offset := tileset.FirstGID - 1
		for i := uint32(0); i < uint32(tileset.TileCount); i++ {
			gid := i + offset
			rect := tileset.GetTileRect(i)
			m.Tiles[gid] = ebiten.NewImageFromImage(tilesetImage.SubImage(rect))
		}
	}
	return nil
}

type Animation struct {
	Tileset   string
	Name      string
	Frames    []*ebiten.Image
	Durations []time.Duration
}

func (m *Tilemap) loadAnimations() {
	var animations []Animation
	for _, tileset := range m.Map.Tilesets {
		for _, tile := range tileset.Tiles {
			if len(tile.Animation) > 0 {
				frames := lo.Map(tile.Animation, func(a *tiled.AnimationFrame, i int) *ebiten.Image {
					return ebiten.NewImageFromImage(m.Tilesets[tileset.Name].SubImage(tileset.GetTileRect(a.TileID)))
				})
				durations := lo.Map(tile.Animation, func(a *tiled.AnimationFrame, i int) time.Duration {
					return time.Millisecond * time.Duration(a.Duration)
				})

				animation := Animation{
					Tileset:   tileset.Name,
					Name:      tile.Properties.GetString("animationName"),
					Durations: durations,
					Frames:    frames,
				}

				animations = append(animations, animation)
			}

		}
	}

	m.Animations = animations
}

func (m *Tilemap) GetAnimation(tileset string, name string) (Animation, bool) {
	return lo.Find(m.Animations, func(a Animation) bool { return a.Tileset == tileset && a.Name == name })
}
