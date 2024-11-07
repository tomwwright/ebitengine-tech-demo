package scenes

import (
	"fmt"
	"image/color"
	"techdemo/components"
	"techdemo/constants"
	"techdemo/events"
	"techdemo/interactions"
	"techdemo/systems"
	"techdemo/tags"
	"techdemo/tilemap"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

const LayerObjects = "Objects"

type TilemapScene struct {
	ecs     *ecs.ECS
	Tilemap *tilemap.Tilemap
	Space   *resolv.Space
}

func NewTilemapScene(filename string) (*TilemapScene, error) {
	tilemap, err := tilemap.LoadTilemap(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to load tilemap from %s: %w", filename, err)
	}

	scene := &TilemapScene{
		ecs:     ecs.NewECS(donburi.NewWorld()),
		Tilemap: tilemap,
	}

	dialogue := systems.NewDialogue()
	playerMovement := systems.NewPlayerMovement()

	debugInputEvents := func(w donburi.World, event events.Input) {
		fmt.Printf("InputEvent: %+v\n", event)
	}

	director := interactions.NewDirector()
	director.SetInteractions(tilemap.Interactions)

	director.RunnableManager.OnStart = func() {
		events.InputEvent.Unsubscribe(scene.ecs.World, playerMovement.OnInputEvent)
		events.InputEvent.Unsubscribe(scene.ecs.World, systems.OnInteractEvent)

	}
	director.RunnableManager.OnFinish = func() {
		events.InputEvent.Subscribe(scene.ecs.World, playerMovement.OnInputEvent)
		events.InputEvent.Subscribe(scene.ecs.World, systems.OnInteractEvent)
	}

	events.InputEvent.Subscribe(scene.ecs.World, debugInputEvents)
	events.InputEvent.Subscribe(scene.ecs.World, playerMovement.OnInputEvent)
	events.InputEvent.Subscribe(scene.ecs.World, systems.OnInteractEvent)
	events.InputEvent.Subscribe(scene.ecs.World, dialogue.OnInteractEvent)

	events.InteractionEvent.Subscribe(scene.ecs.World, director.OnInteractionEvent)

	scene.ecs.AddSystem(systems.NewAnimation().Update)
	scene.ecs.AddSystem(systems.NewMovement().Update)
	scene.ecs.AddSystem(systems.NewInput().Update)
	scene.ecs.AddSystem(systems.ProcessEvents)
	scene.ecs.AddSystem(playerMovement.Update)
	scene.ecs.AddSystem(systems.NewPlayerAnimation(tilemap).Update)
	scene.ecs.AddSystem(systems.NewTextAnimation().Update)
	scene.ecs.AddSystem(systems.UpdateObjects)
	scene.ecs.AddRenderer(ecs.LayerDefault, systems.NewRender().Draw)

	constructSpace(scene)
	constructTileSprites(scene)
	constructObjects(scene)
	constructPlayer(scene)
	constructScreenContainer(scene)
	constructCamera(scene)

	return scene, nil
}

func (s *TilemapScene) Update() {
	s.ecs.Update()
}

func (s *TilemapScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 20, 40, 255})
	s.ecs.Draw(screen)
}

func constructSpace(s *TilemapScene) {
	s.Space = resolv.NewSpace(s.Tilemap.Map.Width*constants.TileSize, s.Tilemap.Map.Height*constants.TileSize, constants.TileSize, constants.TileSize)
}

func constructTileSprites(s *TilemapScene) {
	tilemap := s.Tilemap.Map
	w := s.ecs.World
	for li, l := range tilemap.Layers {
		for i, t := range l.Tiles {
			if !t.Nil {

				entity := w.Create(components.Transform, components.Sprite)
				entry := w.Entry(entity)

				transform := components.Transform.Get(entry)
				xi := i % tilemap.Width
				yi := i / tilemap.Width
				x := float64(l.OffsetX + xi*tilemap.TileWidth)
				y := float64(l.OffsetY + yi*tilemap.TileHeight)
				scale := float64(1)
				transform.LocalPosition = math.NewVec2(x, y)
				transform.LocalScale = math.NewVec2(scale, scale)

				sprite := components.Sprite.Get(entry)
				gid := t.ID + t.Tileset.FirstGID - 1
				sprite.Image = s.Tilemap.Tiles[gid]
				sprite.Layer = li

				if l.Name == LayerObjects {
					object := components.NewObject(entry, constants.Zero, constants.TileSize, constants.TileSize)
					s.Space.Add(&object.Object)
					entry.AddComponent(components.Object)
					components.Object.Set(entry, object)
				}
			}

		}
	}
}

func constructObjects(s *TilemapScene) {
	tilemap := s.Tilemap.Map
	w := s.ecs.World
	for _, group := range tilemap.ObjectGroups {
		for _, o := range group.Objects {
			entity := w.Create(components.Transform, components.Object, components.Interaction)
			entry := w.Entry(entity)

			transform := components.Transform.Get(entry)
			transform.LocalPosition = math.NewVec2(o.X, o.Y)

			object := components.NewObject(entry, constants.Zero, o.Width, o.Height, tags.ResolvTagInteractive)
			s.Space.Add(&object.Object)
			components.Object.Set(entry, object)

			interaction := components.Interaction.Get(entry)
			interaction.Name = o.Name
		}
	}
}

func constructPlayer(s *TilemapScene) {
	w := s.ecs.World
	entity := w.Create(tags.Player, components.Transform, components.Sprite, components.Movement, components.Animation, components.Object)
	entry := w.Entry(entity)

	transform := components.Transform.Get(entry)
	scale := float64(1)
	transform.LocalPosition = math.NewVec2(16, 16)
	transform.LocalScale = math.NewVec2(scale, scale)

	sprite := components.Sprite.Get(entry)
	sprite.Layer = 1

	object := components.NewObject(entry, math.NewVec2(0, 8), constants.TileSize, constants.TileSize/2) // player has collider on lower half of tile only
	s.Space.Add(&object.Object)
	components.Object.Set(entry, object)
}

func constructScreenContainer(s *TilemapScene) {
	w := s.ecs.World
	entity := w.Create(tags.ScreenContainer, components.Transform)
	entry := w.Entry(entity)

	t := components.Transform.Get(entry)
	t.LocalPosition = math.NewVec2(0, 0)

	transform.AppendChild(tags.Player.MustFirst(w), entry, true)
}

func constructCamera(s *TilemapScene) {
	w := s.ecs.World
	entity := w.Create(tags.Camera, components.Transform, components.Movement)
	entry := w.Entry(entity)

	t := components.Transform.Get(entry)
	scale := float64(constants.Scale)
	t.LocalPosition = math.NewVec2(0, 0)
	t.LocalScale = math.NewVec2(scale, scale)
	t.LocalRotation = 0

	transform.AppendChild(tags.ScreenContainer.MustFirst(w), entry, false)
}
