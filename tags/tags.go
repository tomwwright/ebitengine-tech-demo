package tags

import "github.com/yohamta/donburi"

var (
	Player          = donburi.NewTag().SetName("Player")
	Camera          = donburi.NewTag().SetName("Camera")
	ScreenContainer = donburi.NewTag().SetName("ScreenContainer")
	Dialogue        = donburi.NewTag().SetName("Dialogue")
	State           = donburi.NewTag().SetName("State")
	Assets          = donburi.NewTag().SetName("Assets")
	MusicPlayer     = donburi.NewTag().SetName("MusicPlayer")
	FilterChange    = donburi.NewTag().SetName("FilterChange")
)

const ResolvTagInteractive = "Interactive"
const ResolvTagCollider = "Collider"
