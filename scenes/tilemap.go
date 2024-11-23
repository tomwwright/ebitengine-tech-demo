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

	director := scene.Director
	director.RunnableManager.OnStart = func() {
		playerMovement.OnInputEvent(scene.ecs.World, events.InputMoveNone) // cancel player movement
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
	events.TriggerEvent.Subscribe(scene.ecs.World, director.OnTriggerEvent)

	events.MovementFinishedEvent.Subscribe(scene.ecs.World, systems.OnMovementFinishedForTriggers)

	scene.Objects = systems.NewObjects()

	scene.ecs.AddSystem(systems.NewAnimation().Update)
	scene.ecs.AddSystem(systems.NewMovement().Update)
	scene.ecs.AddSystem(systems.NewInput().Update)
	scene.ecs.AddSystem(systems.NewScreenInput().Update)
	scene.ecs.AddSystem(playerMovement.Update)
	scene.ecs.AddSystem(systems.UpdatePlayerAnimation)
	scene.ecs.AddSystem(systems.NewTextAnimation().Update)
	scene.ecs.AddSystem(scene.Objects.Update)
	scene.ecs.AddRenderer(ecs.LayerDefault, systems.NewRender().Draw)

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
	entity := s.ecs.World.Create(tags.State, components.State)
	entry := s.ecs.World.Entry(entity)
	components.State.Set(entry, components.NewState())
}

func constructScreenContainer(s *TilemapScene) {
	w := s.ecs.World
	entity := w.Create(tags.ScreenContainer, components.Transform)
	entry := w.Entry(entity)

	t := components.Transform.Get(entry)
	t.LocalPosition = math.NewVec2(-constants.ScreenWidth/4, -constants.ScreenHeight/4)
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
