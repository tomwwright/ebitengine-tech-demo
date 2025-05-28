package tilemap

import (
	"techdemo/components/collision"
	"techdemo/interactions/yaml"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi/features/math"
)

type TileMap struct {
	Width          int
	Height         int
	Animations     map[string]Animation
	Objects        []Object
	Tiles          []Tile
	TileInstances  []TileInstance
	Interactions   *yaml.Interactions
	PlayerPosition math.Vec2
	PlayerLayer    int
}

type Layer struct {
	DefaultCollision collision.CollisionType
	Tiles            []uint32
}

type TileInstance struct {
	Position  math.Vec2
	Tile      *Tile
	Collision collision.CollisionType
	Layer     int
}

type Object struct {
	Position math.Vec2
	Name     string
}

type Tile struct {
	Name        string
	Animation   Animation
	Collision   collision.CollisionType
	LayerOffset int
}

type Animation struct {
	Name   string
	Frames []Frame
}

func (a Animation) IsAnimated() bool {
	return len(a.Frames) > 1
}

func (a Animation) Image() *ebiten.Image {
	return a.Frames[0].Image
}

type Frame struct {
	Image    *ebiten.Image
	Duration time.Duration
}

func (w *TileMap) GetAnimation(name string) Animation {
	return w.Animations[name]
}
