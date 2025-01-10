package scenes

import (
	"bytes"
	"fmt"
	"techdemo/assets"
	"techdemo/components"
	"techdemo/constants"
	"techdemo/factories"
	"techdemo/systems"
	"techdemo/tags"
	"techdemo/tilemap"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

func LoadTilemapIntoTilemapScene(tilemap *tilemap.Tilemap, scene *TilemapScene) error {

	// configure interactions

	scene.Director.SetInteractions(tilemap.Interactions)

	// spawn tiles and objects

	constructTileSprites(scene, tilemap)
	constructObjects(scene, tilemap)

	// construct space

	scene.Objects.Space = resolv.NewSpace(tilemap.Map.Width*constants.TileSize, tilemap.Map.Height*constants.TileSize, constants.TileSize/2, constants.TileSize/2)

	// attach player to the screen container

	transform.AppendChild(tags.Player.MustFirst(scene.ecs.World), tags.ScreenContainer.MustFirst(scene.ecs.World), false)

	// construct assets

	constructAssets(scene, tilemap)

	// construct music player

	constructMusicPlayer(scene, tilemap)

	return nil
}

const (
	TileClassPlayer = "player"
)

func constructTileSprites(s *TilemapScene, tilemap *tilemap.Tilemap) {
	w := s.ecs.World
	m := tilemap.Map
	for li, l := range m.Layers {

		// get default collision for this layer from properties
		layerCollision := components.CollisionNone
		if c := l.Properties.GetString("collision"); c != "" {
			layerCollision = components.CollisionType(c)
		}

		for i, t := range l.Tiles {
			if !t.Nil {
				xi := i % m.Width
				yi := i / m.Width
				position := math.NewVec2(float64(l.OffsetX+xi*m.TileWidth), float64(l.OffsetY+yi*m.TileHeight))

				collision := layerCollision
				layer := li
				gid := t.ID + t.Tileset.FirstGID - 1
				image := tilemap.Tiles[gid]

				tile, _ := t.Tileset.GetTilesetTile(t.ID)

				if tile != nil {
					if c := tile.Properties.GetString("collision"); c != "" {
						collision = components.CollisionType(c)
					}
					if lo := tile.Properties.GetInt("layerOffset"); lo != 0 {
						layer += lo
					}
				}

				entry := factories.CreateTile(w, t, position, li, collision, image)

				if tile != nil {
					switch tile.Type {
					case TileClassPlayer:
						err := constructPlayer(entry, tilemap)
						if err != nil {
							panic(fmt.Errorf("unable to create player: %w", err))
						}
					}
				}
			}
		}
	}
}

func constructObjects(s *TilemapScene, tilemap *tilemap.Tilemap) {
	m := tilemap.Map
	w := s.ecs.World
	for _, group := range m.ObjectGroups {
		for _, o := range group.Objects {
			entity := w.Create(components.Transform, components.Object, components.Interaction)
			entry := w.Entry(entity)

			transform := components.Transform.Get(entry)
			transform.LocalPosition = math.NewVec2(o.X, o.Y)

			object := components.NewObject(entry, components.CollisionNone, tags.ResolvTagInteractive)
			components.Object.Set(entry, object)

			interaction := components.Interaction.Get(entry)
			interaction.Name = o.Name
		}
	}
}

func constructPlayer(entry *donburi.Entry, tilemap *tilemap.Tilemap) error {

	entry.AddComponent(tags.Player)
	entry.AddComponent(components.Movement)
	entry.AddComponent(components.Animation)
	entry.AddComponent(components.CharacterAnimations)

	animations := components.CharacterAnimations.Get(entry)
	keys := []string{systems.AnimationKeyIdle, systems.AnimationKeyWalkUp, systems.AnimationKeyWalkDown, systems.AnimationKeyWalkRight, systems.AnimationKeyWalkLeft}
	for _, k := range keys {
		a, ok := tilemap.GetAnimation("player", k)
		if !ok {
			return fmt.Errorf("unable to locate player animation: %s", k)
		}
		animations.Add(a)
	}
	return nil
}

func constructAssets(scene *TilemapScene, tilemap *tilemap.Tilemap) {
	w := scene.ecs.World
	entity := w.Create(tags.Assets, components.Assets, components.AudioContext)
	entry := w.Entry(entity)

	components.AudioContext.Set(entry, audio.NewContext(constants.AudioSampleRate))
	components.Assets.Set(entry, &components.AssetsData{
		Assets: assets.NewFileSystemAssets(tilemap.Files),
	})
}

func constructMusicPlayer(scene *TilemapScene, tilemap *tilemap.Tilemap) {
	w := scene.ecs.World
	entity := w.Create(tags.MusicPlayer, components.AudioPlayer)
	entry := w.Entry(entity)

	e := tags.Assets.MustFirst(w)
	asset := components.Assets.Get(e)
	b, _ := asset.Assets.GetAsset(assets.AssetAudioMusic)
	stream, _ := vorbis.DecodeF32(bytes.NewReader(b))

	context := components.AudioContext.Get(e)
	audioPlayer, _ := context.NewPlayerF32(stream)
	audioPlayer.SetVolume(0.3)

	components.AudioPlayer.Set(entry, audioPlayer)

	audioPlayer.Play()
}
