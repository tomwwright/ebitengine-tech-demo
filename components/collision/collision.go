package collision

type CollisionType string

const (
	CollisionUndefined      CollisionType = ""
	CollisionFull           CollisionType = "full"
	CollisionNone           CollisionType = "none"
	CollisionLeft           CollisionType = "left"
	CollisionRight          CollisionType = "right"
	CollisionTop            CollisionType = "top"
	CollisionBottom         CollisionType = "bottom"
	CollisionNotTopLeft     CollisionType = "not_top_left"
	CollisionNotTopRight    CollisionType = "not_top_right"
	CollisionNotBottomLeft  CollisionType = "not_bottom_left"
	CollisionNotBottomRight CollisionType = "not_bottom_right"
)

func (collision CollisionType) Mask() CollisionMask {
	switch collision {
	case CollisionFull:
		return CollisionMask{true, true, true, true}
	case CollisionNone:
		return CollisionMask{false, false, false, false}
	case CollisionLeft:
		return CollisionMask{true, false, true, false}
	case CollisionRight:
		return CollisionMask{false, true, false, true}
	case CollisionTop:
		return CollisionMask{true, true, false, false}
	case CollisionBottom:
		return CollisionMask{false, false, true, true}
	case CollisionNotTopLeft:
		return CollisionMask{false, true, true, true}
	case CollisionNotTopRight:
		return CollisionMask{true, false, true, true}
	case CollisionNotBottomLeft:
		return CollisionMask{true, true, false, true}
	case CollisionNotBottomRight:
		return CollisionMask{true, true, true, false}
	}

	return CollisionMask{false, false, false, false}
}
