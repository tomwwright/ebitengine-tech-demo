package scenes

import (
	"fmt"
	"image/color"
	"io/fs"

	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/factories"
	"github.com/tomwwright/ebitengine-tech-demo/interactions"
	"github.com/tomwwright/ebitengine-tech-demo/observers"
	"github.com/tomwwright/ebitengine-tech-demo/systems"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

type TilemapScene struct {
	ecs      *ecs.ECS
	Director *interactions.Director
	Objects  *systems.ObjectsSystem
}

func NewTilemapScene() (*TilemapScene, error) {
	scene := &TilemapScene{
		ecs:      ecs.NewECS(donburi.NewWorld()),
		Director: interactions.NewDirector(),
	}

	dialogue := systems.NewDialogue()
	playerMovement := systems.NewPlayerMovement()

	debugInputEvents := func(w donburi.World, event events.Input) {
		fmt.Printf("InputEvent: %+v\n", event)
	}

	observers.SetupCurrentInteractionObserver(scene.ecs.World)

	director := scene.Director
	events.InteractionEvent.Subscribe(scene.ecs.World, func(w donburi.World, event events.Interaction) {
		playerMovement.OnInputEvent(scene.ecs.World, events.InputMoveNone) // cancel player movement
		events.InputEvent.Unsubscribe(scene.ecs.World, playerMovement.OnInputEvent)
		events.InputEvent.Unsubscribe(scene.ecs.World, systems.OnInteractEvent)
	})

	events.InteractionFinishedEvent.Subscribe(scene.ecs.World, func(w donburi.World, event events.Interaction) {
		events.InputEvent.Subscribe(scene.ecs.World, playerMovement.OnInputEvent)
		events.InputEvent.Subscribe(scene.ecs.World, systems.OnInteractEvent)
	})

	events.InputEvent.Subscribe(scene.ecs.World, debugInputEvents)
	events.InputEvent.Subscribe(scene.ecs.World, playerMovement.OnInputEvent)
	events.InputEvent.Subscribe(scene.ecs.World, systems.OnInteractEvent)
	events.InputEvent.Subscribe(scene.ecs.World, dialogue.OnInteractEvent)

	events.InteractionEvent.Subscribe(scene.ecs.World, director.OnInteractionEvent)
	events.TriggerEvent.Subscribe(scene.ecs.World, director.OnTriggerEvent)

	events.MovementFinishedEvent.Subscribe(scene.ecs.World, systems.OnMovementFinishedForTriggers)

	scene.Objects = systems.NewObjects()

	render := systems.NewRender()

	scene.ecs.AddSystem(systems.NewAnimation().Update)
	scene.ecs.AddSystem(systems.NewMovement().Update)
	scene.ecs.AddSystem(systems.NewInput().Update)
	scene.ecs.AddSystem(systems.NewScreenInput().Update)
	scene.ecs.AddSystem(playerMovement.Update)
	scene.ecs.AddSystem(systems.UpdatePlayerAnimation)
	scene.ecs.AddSystem(systems.NewTextAnimation().Update)
	scene.ecs.AddSystem(scene.Objects.Update)
	scene.ecs.AddSystem(systems.UpdateFilterChange)
	scene.ecs.AddSystem(systems.UpdateMusicChange)
	scene.ecs.AddSystem(systems.UpdateZoomChange)
	scene.ecs.AddSystem(render.Update)
	scene.ecs.AddRenderer(ecs.LayerDefault, render.Draw)

	// process events as last system to ensure all component data has been updated
	scene.ecs.AddSystem(systems.ProcessEvents)

	constructState(scene)
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

func constructState(s *TilemapScene) {
	e := archetypes.State.Create(s.ecs.World)
	components.State.Set(e, components.NewState())
}

func constructScreenContainer(s *TilemapScene) {
	world := s.ecs.World
	e := archetypes.ScreenContainer.Create(world)
	transform.GetTransform(e).LocalScale = math.NewVec2(constants.Scale, constants.Scale)
}

func constructCamera(s *TilemapScene) {
	w := s.ecs.World
	e := factories.CreateCamera(w)
	container := archetypes.CameraContainer.Create(w)
	transform.AppendChild(container, e, false)
}

func (s *TilemapScene) ConfigureAssets(files fs.ReadFileFS) {
	factories.CreateAssets(s.ecs.World, files)
}
