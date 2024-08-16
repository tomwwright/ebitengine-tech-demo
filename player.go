package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tanema/gween/ease"
)

type Player struct {
	Position      [2]float32
	Sprite        *ebiten.Image
	MoveAnimation *SliceTween
}

func (p *Player) Update() {
	if p.MoveAnimation != nil {
		position, isFinished := p.MoveAnimation.Update(dt)
		p.Position = [2]float32(position)
		if isFinished {
			p.MoveAnimation = nil
		}
	}

	if p.MoveAnimation == nil {
		d := float32(16)
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.setMoveAnimation([2]float32{p.Position[0] + d, p.Position[1]})
		} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.setMoveAnimation([2]float32{p.Position[0] - d, p.Position[1]})
		} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.setMoveAnimation([2]float32{p.Position[0], p.Position[1] - d})
		} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.setMoveAnimation([2]float32{p.Position[0], p.Position[1] + d})
		}
	}

}

func (p *Player) setMoveAnimation(to [2]float32) {
	p.MoveAnimation = NewSliceTween(p.Position[:], to[:], 0.4, ease.Linear)
}

func (p *Player) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(
		float64(p.Position[0]),
		float64(p.Position[1]),
	)
	op.GeoM.Scale(scale, scale)

	screen.DrawImage(p.Sprite, op)
}
