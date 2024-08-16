package main

import (
	"bytes"
	"fmt"
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

const (
	screenWidth  = 640
	screenHeight = 480
	scale        = 2
	dt           = 1.0 / 60.0
)

type Game struct {
	Map    *tiled.Map
	Player *Player
	Tiles  map[uint32]*ebiten.Image
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, l := range g.Map.Layers {
		for i, t := range l.Tiles {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(l.OffsetX+(i%g.Map.Width)*g.Map.TileWidth),
				float64(l.OffsetY+(i/g.Map.Width)*g.Map.TileHeight),
			)
			op.GeoM.Scale(scale, scale)

			if !t.Nil {
				gid := t.ID + t.Tileset.FirstGID - 1
				tile := g.Tiles[gid]
				screen.DrawImage(tile, op)
			}

		}
	}

	g.Player.Draw(screen)

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) loadTilesets() error {
	g.Tiles = map[uint32]*ebiten.Image{}
	for _, tileset := range g.Map.Tilesets {
		fmt.Println("loading tileset", tileset.Name)
		source, err := os.ReadFile("tilesets/" + tileset.Image.Source)
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
			g.Tiles[gid] = ebiten.NewImageFromImage(tilesetImage.SubImage(rect))
		}
	}
	return nil
}

func main() {

	gameMap, err := tiled.LoadFile("tilesets/tilemap.tmx")
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	game := Game{
		Map: gameMap,
	}

	err = game.loadTilesets()
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	game.Player = &Player{
		Position: [2]float32{16, 16},
		Sprite:   game.Tiles[217],
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
