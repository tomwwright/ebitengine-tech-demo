package main

import (
	"embed"
	"fmt"
	_ "image/png"
	"log"
	"techdemo/constants"
	"techdemo/scenes"

	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed tilesets/*
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

	scene, err := scenes.NewTilemapScene(files, "tilesets/tilemap.tmx")
	if err != nil {
		panic(fmt.Sprintf("failed to load scene: %v", err))
	}

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
