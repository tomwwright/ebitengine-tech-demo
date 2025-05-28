package tilemap

import (
	"techdemo/components/collision"
)

func BuildFromTiled(tiled Tiled) *TileMap {
	builder := &builder{
		tiled,
	}
	return builder.Build()
}

type builder struct {
	Tiled
}

const (
	TilePlayer = "player"
)

func (b *builder) Build() *TileMap {
	width, height := b.Size()
	world := &TileMap{
		Width:         width,
		Height:        height,
		Objects:       b.Objects(),
		Tiles:         b.Tiles(),
		TileInstances: []TileInstance{},
		Interactions:  b.Interactions(),
		Animations:    map[string]Animation{},
	}

	b.buildAnimations(world)
	b.buildTileInstances(world)

	return world
}

func (b *builder) buildAnimations(w *TileMap) {
	for _, a := range b.Animations() {
		w.Animations[a.Name] = a
	}
}

func (b *builder) buildTileInstances(w *TileMap) {
	for li, l := range b.Tiled.Layers() {

		for i, id := range l.Tiles {
			if id != 0 {

				t := w.Tiles[id-1]

				position := b.ToPixels(b.ToPosition(i))

				// save player location if found
				if t.Name == TilePlayer {
					w.PlayerPosition = position
					w.PlayerLayer = li
					continue
				}

				// default tile

				collisionType := l.DefaultCollision
				if t.Collision != collision.CollisionUndefined {
					collisionType = t.Collision
				}
				layer := li + t.LayerOffset

				w.TileInstances = append(w.TileInstances, TileInstance{
					Position:  position,
					Layer:     layer,
					Collision: collisionType,
					Tile:      &t,
				})
			}
		}
	}
}
