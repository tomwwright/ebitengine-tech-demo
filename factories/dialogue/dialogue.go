package dialogue

import (
	"bytes"
	"image/color"
	"strings"

	"github.com/tomwwright/ebitengine-tech-demo/archetypes"
	"github.com/tomwwright/ebitengine-tech-demo/assets"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/events"
	"github.com/tomwwright/ebitengine-tech-demo/tags"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/features/transform"
)

const DialogueSpeed = 12.0
const FontWidth = 6.0
const FontScaling = 0.75
const DialogueMaxLineLength = (constants.ScreenWidth - constants.TileSize*constants.Scale) / (FontWidth * constants.Scale * FontScaling)

var BackdropImage = createDialogueBackdrop()
var DialogueFont = text.NewGoXFace(bitmapfont.Face)

func CreateDialogue(w donburi.World, text string) {
	CloseDialogue(w)

	text = insertLineBreaks(text)

	// backdrop

	backdrop := archetypes.Dialogue.Create(w)

	t := components.Transform.Get(backdrop)
	t.LocalPosition = math.NewVec2(constants.TileSize, float64(constants.ScreenHeight-BackdropImage.Bounds().Dy()*constants.Scale-constants.TileSize))
	components.Sprite.Set(backdrop, &components.SpriteData{
		Image: BackdropImage,
		Layer: constants.LayerUI,
	})

	transform.AppendChild(tags.ScreenContainer.MustFirst(w), backdrop, false)

	// audio

	e := tags.Assets.MustFirst(w)
	asset := components.Assets.Get(e)
	b, _ := asset.Assets.GetAsset(assets.AssetAudioText)
	stream, _ := wav.DecodeF32(bytes.NewReader(b))

	context := components.AudioContext.Get(e)
	audioPlayer, _ := context.NewPlayerF32(stream)
	audioPlayer.SetVolume(0.4)

	components.AudioPlayer.Set(backdrop, audioPlayer)

	playTextScroll := func(n int) {
		isWhitespace := n > len(text)-1 || text[n] == ' ' || text[n] == '\n'
		if !isWhitespace {
			audioPlayer.SetPosition(0)
			audioPlayer.Play()
		}
	}

	// text

	entry := archetypes.Text.Create(w)

	t = components.Transform.Get(entry)
	t.LocalPosition = math.NewVec2(constants.TileSize, constants.TileSize)
	t.LocalScale = math.NewVec2(FontScaling, FontScaling)
	components.Text.Set(entry, &components.TextData{
		Font:  DialogueFont,
		Text:  "",
		Layer: constants.LayerUI,
	})
	components.TextAnimation.Set(entry, &components.TextAnimationData{
		Speed:  DialogueSpeed,
		Text:   text,
		OnTick: playTextScroll,
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
	w := constants.Width - constants.TileSize/2
	h := constants.Height / 3
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
		if float32(length) <= DialogueMaxLineLength {
			text += " " + w
		} else {
			text += "\n" + w
			length = len(w)
		}
	}
	return text
}
