package scenes

import (
	"fmt"

	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
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
		e := archetypes.Interaction.Create(w)

		transform := components.Transform.Get(e)
		transform.LocalPosition = o.Position

		object := components.NewObject(e, constants.TileSize, constants.TileSize, tags.ResolvTagInteractive)
		components.Object.Set(e, object)

		interaction := components.Interaction.Get(e)
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

	e, err := factories.CreateMusicPlayer(w, assets.AssetAudioMusic)
	if err != nil {
		return fmt.Errorf("failed to create music: %w", err)
	}
	audio := components.AudioPlayer.Get(e)
	audio.SetVolume(0.3)
	audio.Play()

	return nil
}
