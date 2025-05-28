package tiled

import (
	"image"
	"time"

	"github.com/tomwwright/ebitengine-tech-demo/components/collision"
	"github.com/tomwwright/ebitengine-tech-demo/interactions/yaml"
	"github.com/tomwwright/ebitengine-tech-demo/tilemap"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"github.com/samber/lo"
	"github.com/yohamta/donburi/features/math"
)

type Tiled struct {
	tilemap      *tiled.Map
	tilesets     map[string]image.Image
	interactions *yaml.Interactions
}

func (t *Tiled) Size() (int, int) {
	return t.tilemap.Width, t.tilemap.Height
}

func (t *Tiled) ToPosition(index int) math.Vec2 {
	xi := index % t.tilemap.Width
	yi := index / t.tilemap.Width
	return math.NewVec2(float64(xi), float64(yi))
}

func (t *Tiled) ToPixels(position math.Vec2) math.Vec2 {
	return position.Mul(math.NewVec2(float64(t.tilemap.TileWidth), float64(t.tilemap.TileHeight)))
}

func (t *Tiled) Interactions() *yaml.Interactions {
	return t.interactions
}

func (t *Tiled) Tiles() (tiles []tilemap.Tile) {

	n := 0
	for _, tileset := range t.tilemap.Tilesets {
		n += tileset.TileCount
	}
	tiles = make([]tilemap.Tile, n)

	for _, tileset := range t.tilemap.Tilesets {

		for i := range uint32(tileset.TileCount) {
			gid := tileset.FirstGID + uint32(i) - 1
			tiledTile, _ := tileset.GetTilesetTile(i)

			var tile tilemap.Tile
			if tiledTile == nil {
				img := ebiten.NewImageFromImage(ebiten.NewImageFromImage(t.tilesets[tileset.Name]).SubImage(tileset.GetTileRect(i)))
				tile = tilemap.Tile{
					Animation: tilemap.Animation{
						Frames: []tilemap.Frame{
							tilemap.Frame{
								Image: img,
							},
						},
					},
				}
			} else {
				collisionType := collision.CollisionUndefined
				if c := tiledTile.Properties.GetString("collision"); c != "" {
					collisionType = collision.CollisionType(c)
				}

				layerOffset := 0
				if lo := tiledTile.Properties.GetInt("layerOffset"); lo != 0 {
					layerOffset = lo
				}

				frames := []tilemap.Frame{}
				if tiledTile.Animation != nil {
					for _, frame := range tiledTile.Animation {

						img := ebiten.NewImageFromImage(ebiten.NewImageFromImage(t.tilesets[tileset.Name]).SubImage(tileset.GetTileRect(frame.TileID)))
						duration := time.Duration(frame.Duration * uint32(time.Millisecond))

						frames = append(frames, tilemap.Frame{
							Image:    img,
							Duration: duration,
						})
					}
				} else {
					img := ebiten.NewImageFromImage(ebiten.NewImageFromImage(t.tilesets[tileset.Name]).SubImage(tileset.GetTileRect(tiledTile.ID)))
					frames = append(frames, tilemap.Frame{
						Image: img,
					})
				}

				tile = tilemap.Tile{
					Name: tiledTile.Type,
					Animation: tilemap.Animation{
						Frames: frames,
					},
					Collision:   collisionType,
					LayerOffset: layerOffset,
				}
			}

			tiles[gid] = tile
		}
	}
	return tiles
}

func (t *Tiled) Animations() (animations []tilemap.Animation) {
	animations = []tilemap.Animation{}
	for _, tileset := range t.tilemap.Tilesets {
		for _, tile := range tileset.Tiles {
			if tile != nil && tile.Animation != nil {
				name := tile.Properties.GetString("animationName")
				if name != "" {
					frames := []tilemap.Frame{}
					for _, frame := range tile.Animation {
						img := ebiten.NewImageFromImage(ebiten.NewImageFromImage(t.tilesets[tileset.Name]).SubImage(tileset.GetTileRect(frame.TileID)))
						duration := time.Duration(frame.Duration * uint32(time.Millisecond))

						frames = append(frames, tilemap.Frame{
							Image:    img,
							Duration: duration,
						})
					}
					animations = append(animations, tilemap.Animation{
						Name:   tileset.Name + "/" + name,
						Frames: frames,
					})
				}
			}
		}
	}
	return animations
}

func (t *Tiled) Objects() []tilemap.Object {
	objects := []tilemap.Object{}
	for _, group := range t.tilemap.ObjectGroups {
		for _, o := range group.Objects {

			object := tilemap.Object{
				Position: math.NewVec2(o.X, o.Y),
				Name:     o.Name,
			}

			objects = append(objects, object)
		}
	}
	return objects
}

func (t *Tiled) Layers() []tilemap.Layer {
	layers := []tilemap.Layer{}
	for _, l := range t.tilemap.Layers {
		if !l.Visible {
			continue
		}

		defaultCollision := collision.CollisionNone
		if c := l.Properties.GetString("collision"); collision.CollisionType(c) != collision.CollisionUndefined {
			defaultCollision = collision.CollisionType(c)
		}

		getTileID := func(tile *tiled.LayerTile) uint32 {
			if tile.Nil {
				return 0
			} else {
				return tile.Tileset.FirstGID + tile.ID
			}
		}

		tiles := lo.Map(l.Tiles, func(tile *tiled.LayerTile, i int) uint32 { return getTileID(tile) })

		layers = append(layers, tilemap.Layer{
			DefaultCollision: defaultCollision,
			Tiles:            tiles,
		})
	}
	return layers
}
