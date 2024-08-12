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
)

type Game struct {
	Map     *tiled.Map
	Tileset *ebiten.Image
	Tiles   map[uint32]*ebiten.Image
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	tileset := g.Map.Tilesets[0]
	sx := float64(screenWidth) / float64(g.Map.Width*tileset.TileWidth)
	sy := float64(screenHeight) / float64(g.Map.Height*tileset.TileHeight)

	for _, l := range g.Map.Layers {
		for i, t := range l.Tiles {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(
				float64(l.OffsetX+(i%g.Map.Width)*g.Map.TileWidth),
				float64(l.OffsetY+(i/g.Map.Width)*g.Map.TileHeight),
			)
			op.GeoM.Scale(sx, sy)

			id := t.ID
			if !t.Nil {
				tile := g.Tiles[id]
				screen.DrawImage(tile, op)
			}

		}
	}

	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) prepareTiles() {
	tileset := g.Map.Tilesets[0]
	g.Tiles = map[uint32]*ebiten.Image{}
	for i := uint32(0); i < uint32(tileset.TileCount); i++ {
		rect := tileset.GetTileRect(i)
		g.Tiles[i] = ebiten.NewImageFromImage(g.Tileset.SubImage(rect))
	}
}

func main() {

	// Parse .tmx file.
	gameMap, err := tiled.LoadFile("untitled.tmx")
	if err != nil {
		fmt.Printf("error parsing map: %s", err.Error())
		os.Exit(2)
	}

	tileset := gameMap.Tilesets[0]
	source, err := os.ReadFile(tileset.Image.Source)

	tilesetImage, _, err := image.Decode(bytes.NewReader(source))
	if err != nil {
		fmt.Printf("error loading tileset: %s", err.Error())
		os.Exit(2)
	}

	game := Game{
		Map:     gameMap,
		Tileset: ebiten.NewImageFromImage(tilesetImage),
	}

	game.prepareTiles()

	fmt.Println("drawing with scale", float64(screenWidth)/float64(game.Map.Width*tileset.TileWidth))

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
