package tilemap

import (
	"bytes"
	"fmt"
	"image"
	"io/fs"
	"path/filepath"
	"techdemo/interactions/yaml"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/samber/lo"
)

type Tilemap struct {
	Map          *tiled.Map
	Filename     string
	Tiles        map[uint32]*ebiten.Image
	Tilesets     map[string]*ebiten.Image
	Animations   []Animation
	Interactions *yaml.Interactions
	Files        fs.ReadFileFS
}

func LoadTilemap(files fs.ReadFileFS, filename string) (*Tilemap, error) {
	m := &Tilemap{
		Filename: filename,
		Files:    files,
	}
	err := m.loadMap()
	if err != nil {
		return nil, err
	}

	err = m.loadTilesets()
	if err != nil {
		return nil, err
	}

	err = m.loadInteractions()
	if err != nil {
		return nil, err
	}

	m.loadAnimations()

	return m, nil
}

func (m *Tilemap) loadMap() error {
	tiledMap, err := tiled.LoadFile(m.Filename, tiled.WithFileSystem(m.Files))
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
		filename := dir + "/" + tileset.Image.Source
		fmt.Println("loading tileset", tileset.Name, "from", filename)
		source, err := m.Files.ReadFile(filename)
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

func (m *Tilemap) loadInteractions() error {
	filename := m.Map.Properties.Get("interactionsFilename")[0]
	dir := filepath.Dir(m.Filename)
	filepath := dir + "/" + filename
	content, err := m.Files.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to load interactions from %s: %w", filepath, err)
	}
	interactions, err := yaml.UnmarshallInteractions(content)
	if err != nil {
		return err
	}
	m.Interactions = interactions
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

				name := tile.Properties.GetString("animationName")
				if name == "" {
					name = fmt.Sprintf("Animated Tile %d (Tileset %s)", tile.ID, tileset.Name)
				}

				animation := Animation{
					Tileset:   tileset.Name,
					Name:      name,
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
