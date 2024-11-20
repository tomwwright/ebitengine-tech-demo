package tags

import "github.com/yohamta/donburi"

var (
	Player          = donburi.NewTag().SetName("Player")
	Camera          = donburi.NewTag().SetName("Camera")
	ScreenContainer = donburi.NewTag().SetName("ScreenContainer")
	Dialogue        = donburi.NewTag().SetName("Dialogue")
	State           = donburi.NewTag().SetName("State")
)

const ResolvTagInteractive = "Interactive"
const ResolvTagCollider = "Collider"
