package main

import (
	"embed"
	"fmt"
	_ "image/png"
	"log"
	"techdemo/constants"
	"techdemo/scenes"
	"techdemo/tiled"
	"techdemo/tilemap"

	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed tilesets assets
var files embed.FS

type Scene interface {
	Update()
	Draw(screen *ebiten.Image)
}

type Game struct {
	bounds image.Rectangle
	scene  Scene
}

func NewGame() *Game {

	dir, _ := files.ReadDir("tilesets")
	fmt.Printf("%+v\n", dir)

	filename := "tilesets/new.tmx"

	t, err := tiled.Load(filename, tiled.WithFiles(files))
	if err != nil {
		panic(fmt.Errorf("failed to load tilemap from %s: %w", filename, err))
	}

	world := tilemap.BuildFromTiled(t)

	scene, err := scenes.NewTilemapScene()
	if err != nil {
		panic(fmt.Sprintf("failed to create scene: %v", err))
	}

	scene.ConfigureAssets(files)

	err = scenes.LoadScene(world, scene)
	if err != nil {
		panic(fmt.Errorf("failed to load scene: %w", filename, err))
	}

	// tilemap, err := tilemap.LoadTilemap(files, filename)
	// if err != nil {
	// 	panic(fmt.Errorf("failed to load tilemap from %s: %w", filename, err))
	// }

	// scene, err = scenes.NewTilemapScene()
	// if err != nil {
	// 	panic(fmt.Sprintf("failed to load scene: %v", err))
	// }

	// scenes.LoadTilemapIntoTilemapScene(tilemap, scene)

	g := &Game{
		bounds: image.Rectangle{},
		scene:  scene,
	}

	return g
}

func (g *Game) Update() error {
	g.scene.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Clear()
	g.scene.Draw(screen)
}

func (g *Game) Layout(width, height int) (int, int) {
	g.bounds = image.Rect(0, 0, width, height)
	return width, height
}

func main() {
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
