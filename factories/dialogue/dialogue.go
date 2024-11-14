package dialogue

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

var BackdropImage = createDialogueBackdrop()
var DialogueFont = text.NewGoXFace(bitmapfont.Face)

func CreateDialogue(w donburi.World, text string) {
	CloseDialogue(w)

	text = insertLineBreaks(text)

	// backdrop

	entity := w.Create(tags.Dialogue, components.Transform, components.Sprite)
	backdrop := w.Entry(entity)

	t := components.Transform.Get(backdrop)
	t.LocalPosition = math.NewVec2(constants.TileSize/2, float64(constants.ScreenHeight/constants.Scale-BackdropImage.Bounds().Dy()-constants.TileSize/2))
	components.Sprite.Set(backdrop, &components.SpriteData{
		Image: BackdropImage,
		Layer: constants.LayerUI,
	})

	transform.AppendChild(tags.ScreenContainer.MustFirst(w), backdrop, false)

	// text

	entity = w.Create(components.Transform, components.Text, components.TextAnimation)
	entry := w.Entry(entity)

	t = components.Transform.Get(entry)
	t.LocalPosition = math.NewVec2(constants.TileSize/2, constants.TileSize/2)
	components.Text.Set(entry, &components.TextData{
		Font: DialogueFont,
		Text: "",
	})
	components.TextAnimation.Set(entry, &components.TextAnimationData{
		Speed: DialogueSpeed,
		Text:  text,
	})

	transform.AppendChild(backdrop, entry, false)
	events.StateChangeEvent.Publish(w, events.DialogueOpened)
}

func CloseDialogue(w donburi.World) {
	existingDialogue, _ := tags.Dialogue.First(w)
	if existingDialogue != nil {
		transform.RemoveRecursive(existingDialogue)
		events.StateChangeEvent.Publish(w, events.DialogueClosed)
	}
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
