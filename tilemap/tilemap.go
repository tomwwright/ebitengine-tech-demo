package tilemap

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type Tilemap struct {
	Map      *tiled.Map
	Filename string
	Tiles    map[uint32]*ebiten.Image
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

		offset := tileset.FirstGID - 1
		for i := uint32(0); i < uint32(tileset.TileCount); i++ {
			gid := i + offset
			fmt.Println("loading tile", i, "global", gid)
			rect := tileset.GetTileRect(i)
			m.Tiles[gid] = ebiten.NewImageFromImage(tilesetImage.SubImage(rect))
		}
	}
	return nil
}
