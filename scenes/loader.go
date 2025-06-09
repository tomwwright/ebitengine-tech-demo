package scenes

import (
	"bytes"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/factories"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/tomwwright/ebitengine-tech-demo/tilemap"

	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi/features/transform"
)

func LoadScene(world *tilemap.TileMap, scene *TilemapScene) error {
	w := scene.ecs.World

	// configure interactions

	scene.Director.SetInteractions(world.Interactions)

	// spawn tiles

	for _, instance := range world.TileInstances {
		entry := factories.CreateTile(w, instance.Position, instance.Layer, instance.Collision, instance.Tile.Animation.Image())
		if instance.Tile.Animation.IsAnimated() {
			entry.AddComponent(components.Animation)
			data := components.NewAnimationFromTilemapAnimation(instance.Tile.Animation)
			components.Animation.Set(entry, data)
		}
	}

	// spawn objects

	for _, o := range world.Objects {
		entity := w.Create(components.Transform, components.Object, components.Interaction)
		entry := w.Entry(entity)

		transform := components.Transform.Get(entry)
		transform.LocalPosition = o.Position

		object := components.NewObject(entry, constants.TileSize, constants.TileSize, tags.ResolvTagInteractive)
		components.Object.Set(entry, object)

		interaction := components.Interaction.Get(entry)
		interaction.Name = o.Name
	}

	// spawn player

	_, err := factories.CreatePlayer(w, world, world.PlayerPosition, world.PlayerLayer)
	if err != nil {
		return fmt.Errorf("failed to create player: %w", err)
	}

	// construct space

	scene.Objects.Space = resolv.NewSpace(world.Width*constants.TileSize, world.Height*constants.TileSize, constants.TileSize/2, constants.TileSize/2)

	// attach player to the screen container

	transform.AppendChild(tags.Player.MustFirst(scene.ecs.World), tags.ScreenContainer.MustFirst(scene.ecs.World), false)

	// construct music player

	e := tags.Assets.MustFirst(w)

	context := components.AudioContext.Get(e)
	asset := components.Assets.Get(e)

	entity := w.Create(tags.MusicPlayer, components.AudioPlayer, components.Asset)
	entry := w.Entry(entity)

	components.Asset.SetValue(entry, components.AssetData{
		Asset: assets.AssetAudioMusic,
	})
	b, _ := asset.Assets.GetAsset(assets.AssetAudioMusic)
	stream, _ := vorbis.DecodeF32(bytes.NewReader(b))
	audio, _ := context.NewPlayerF32(stream)
	audio.SetVolume(0.3)
	audio.Play()
	components.AudioPlayer.Set(entry, audio)

	return nil
}
