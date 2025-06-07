package components

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type CameraData struct {
	ToScreen ebiten.GeoM
	Viewport math.Vec2
	Color    color.Color
}

var Camera = donburi.NewComponentType[CameraData]()

func (c *CameraData) Calculate(position math.Vec2, scale math.Vec2, rotation float64) {
	m := ebiten.GeoM{}

	viewportCenter := c.Viewport.MulScalar(0.5)

	m.Translate(-position.X, -position.Y)

	m.Translate(-viewportCenter.X/scale.X, -viewportCenter.Y/scale.Y) // rotate around center of viewport, considering scale
	m.Rotate(rotation)

	m.Translate(viewportCenter.X/scale.X, viewportCenter.Y/scale.Y)

	m.Scale(
		scale.X,
		scale.Y,
	)

	c.ToScreen = m
}

func (c *CameraData) SetViewportFromImage(image *ebiten.Image) {
	c.Viewport.X = float64(image.Bounds().Dx())
	c.Viewport.Y = float64(image.Bounds().Dy())
}

func (c *CameraData) IsVisible(position math.Vec2, size math.Vec2) bool {
	minX, minY := c.ToScreen.Apply(position.XY())
	maxX, maxY := c.ToScreen.Apply(position.Add(size).XY())
	if maxX < -c.Viewport.X*0.1 || maxY < -c.Viewport.Y*0.1 || minX > c.Viewport.X*1.1 || minY > c.Viewport.Y*1.1 {
		return false
	}
	return true
}

func (c *CameraData) Draw(image *ebiten.Image, position math.Vec2, scale math.Vec2, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleWithColor(c.Color)
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)
	op.GeoM.Concat(c.ToScreen)
	screen.DrawImage(image, op)
}

func (c *CameraData) DrawText(font text.Face, message string, position math.Vec2, scale math.Vec2, screen *ebiten.Image) {
	op := &text.DrawOptions{
		LayoutOptions: text.LayoutOptions{
			LineSpacing: constants.LineSpacing,
		},
	}
	op.GeoM.Scale(scale.X, scale.Y)
	op.GeoM.Translate(position.X, position.Y)
	op.GeoM.Concat(c.ToScreen)
	text.Draw(screen, message, font, op)
}
