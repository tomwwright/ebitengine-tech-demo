package systems

import (
	"image/color"
	"strings"
	"techdemo/components"
	"techdemo/constants"
	"techdemo/events"
	"techdemo/tags"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

const DialogueSpeed = 12.0
const DialogueMaxLineLength = 48

type Dialogue struct {
	backdrop *ebiten.Image
	font     text.Face
	current  *currentDialogue
}

type currentDialogue struct {
	backdrop *donburi.Entry
	text     *donburi.Entry
}

func NewDialogue() *Dialogue {
	return &Dialogue{
		backdrop: createDialogueBackdrop(),
		font:     text.NewGoXFace(bitmapfont.Face),
	}
}

func (d *Dialogue) OnDialogueEvent(w donburi.World, event events.Dialogue) {

	if d.current != nil {
		return
	}

	text := insertLineBreaks(event.Text)
	d.SetDialogue(w, text)
	events.StateChangeEvent.Publish(w, events.DialogueOpened)
}

func (d *Dialogue) SetDialogue(w donburi.World, text string) {
	if d.current != nil {
		d.closeCurrent()
	}

	// backdrop

	entity := w.Create(components.Transform, components.Sprite)
	backdrop := w.Entry(entity)

	t := components.Transform.Get(backdrop)
	t.LocalPosition = math.NewVec2(constants.TileSize/2, float64(constants.ScreenHeight/constants.Scale-d.backdrop.Bounds().Dy()-constants.TileSize/2))
	components.Sprite.Set(backdrop, &components.SpriteData{
		Image: d.backdrop,
		Layer: constants.LayerUI,
	})

	transform.AppendChild(tags.ScreenContainer.MustFirst(w), backdrop, false)

	// text

	entity = w.Create(components.Transform, components.Text, components.TextAnimation)
	entry := w.Entry(entity)

	t = components.Transform.Get(entry)
	t.LocalPosition = math.NewVec2(constants.TileSize/2, constants.TileSize/2)
	components.Text.Set(entry, &components.TextData{
		Font: d.font,
		Text: "",
	})
	components.TextAnimation.Set(entry, &components.TextAnimationData{
		Speed: DialogueSpeed,
		Text:  text,
	})

	transform.AppendChild(backdrop, entry, false)

	d.current = &currentDialogue{
		backdrop: backdrop,
		text:     entry,
	}
}

func (d *Dialogue) OnInteractEvent(w donburi.World, event events.Input) {
	if d.current == nil || event != events.InputInteract {
		return
	}

	animation := components.TextAnimation.Get(d.current.text)
	if animation.IsFinished() {
		d.closeCurrent()
		events.StateChangeEvent.Publish(w, events.DialogueClosed)
	} else {
		animation.Skip()
	}
}

func (d *Dialogue) closeCurrent() {
	if d.current == nil {
		return
	}

	d.current.text.Remove()
	d.current.backdrop.Remove()
	d.current = nil
}

func createDialogueBackdrop() *ebiten.Image {
	w := constants.ScreenWidth/constants.Scale - constants.TileSize
	h := constants.ScreenHeight/constants.Scale/3 - constants.TileSize/2
	img := ebiten.NewImage(w, h)
	img.Fill(color.RGBA{0, 0, 0, 100})
	return img
}

func insertLineBreaks(text string) string {
	words := strings.Split(text, " ")
	if len(words) < 2 {
		return text
	}

	text = words[0]
	length := len(text)
	for _, w := range words[1:] {
		length += len(w) + 1 // for space
		if length <= DialogueMaxLineLength {
			text += " " + w
		} else {
			text += "\n" + w
			length = len(w)
		}
	}
	return text
}
