package scenes

import (
	"fmt"

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
		factories.CreateObject(w, o.Position, o.Size, o.Name)
	}

	// spawn player

	_, err := factories.CreatePlayer(w, world, world.PlayerPosition, world.PlayerLayer)
	if err != nil {
		return fmt.Errorf("failed to create player: %w", err)
	}

	// construct space

	scene.Objects.Space = resolv.NewSpace(world.Width*constants.TileSize, world.Height*constants.TileSize, constants.TileSize/2, constants.TileSize/2)

	// attach camera container to the player

	e := tags.CameraContainer.MustFirst(scene.ecs.World)
	t := transform.GetTransform(e)
	t.LocalPosition.X = constants.TileSize / 2
	t.LocalPosition.Y = constants.Width / 3

	transform.AppendChild(tags.Player.MustFirst(scene.ecs.World), e, false)

	// construct music player

	e, err = factories.CreateMusicPlayer(w, assets.AssetAudioMusic)
	if err != nil {
		return fmt.Errorf("failed to create music: %w", err)
	}
	audio := components.AudioPlayer.Get(e)
	audio.SetVolume(0.3)
	audio.Play()

	return nil
}
