package collision

import (
	"github.com/yohamta/donburi/features/math"
)

type CollisionMask struct {
	TopLeft     bool
	TopRight    bool
	BottomLeft  bool
	BottomRight bool
}

type CollisionObject struct {
	Position math.Vec2
	Size     math.Vec2
}

func (m CollisionMask) Objects() []CollisionObject {
	n := m.count()
	if n == 4 {
		// 0 0
		// 0 0
		return []CollisionObject{{math.Vec2{X: 0, Y: 0}, math.Vec2{X: 1, Y: 1}}}
	} else if n == 3 {
		if !m.TopLeft {
			// X 1
			// 0 0
			return []CollisionObject{
				{
					Position: math.Vec2{X: 0, Y: 0.5},
					Size:     math.Vec2{X: 1, Y: 0.5},
				},
				{
					Position: math.Vec2{X: 0.5, Y: 0},
					Size:     math.Vec2{X: 0.5, Y: 0.5},
				},
			}
		}
		if !m.TopRight {
			// 0 X
			// 1 1
			return []CollisionObject{
				{
					Position: math.Vec2{X: 0, Y: 0},
					Size:     math.Vec2{X: 0.5, Y: 0.5},
				},
				{
					Position: math.Vec2{X: 0, Y: 0.5},
					Size:     math.Vec2{X: 1, Y: 0.5},
				},
			}
		}
		if !m.BottomLeft {
			// 0 0
			// X 1
			return []CollisionObject{
				{
					Position: math.Vec2{X: 0, Y: 0},
					Size:     math.Vec2{X: 1, Y: 0.5},
				},
				{
					Position: math.Vec2{X: 0.5, Y: 0.5},
					Size:     math.Vec2{X: 0.5, Y: 0.5},
				},
			}
		}
		if !m.BottomRight {
			// 0 0
			// 1 X
			return []CollisionObject{
				{
					Position: math.Vec2{X: 0, Y: 0},
					Size:     math.Vec2{X: 1, Y: 0.5},
				},
				{
					Position: math.Vec2{X: 0, Y: 0.5},
					Size:     math.Vec2{X: 0.5, Y: 0.5},
				},
			}
		}
	} else if n == 2 {
		if m.Top() {
			// 0 0
			// X X
			return []CollisionObject{{math.Vec2{X: 0, Y: 0}, math.Vec2{X: 1, Y: 0.5}}}
		}
		if m.Bottom() {
			// X X
			// 0 0
			return []CollisionObject{{math.Vec2{X: 0, Y: 0.5}, math.Vec2{X: 1, Y: 0.5}}}
		}
		if m.Left() {
			// 0 X
			// 0 X
			return []CollisionObject{{math.Vec2{X: 0, Y: 0}, math.Vec2{X: 0.5, Y: 1}}}
		}
		if m.Right() {
			// X 0
			// X 0
			return []CollisionObject{{math.Vec2{X: 0.5, Y: 0}, math.Vec2{X: 0.5, Y: 1}}}
		}
		return nil // diagonals unsupported
	} else if n == 1 {
		return nil // corners unsupported
	}
	return nil
}

func (m CollisionMask) Top() bool {
	return m.TopLeft && m.TopRight
}

func (m CollisionMask) Bottom() bool {
	return m.BottomLeft && m.BottomRight
}

func (m CollisionMask) Left() bool {
	return m.TopLeft && m.BottomLeft
}

func (m CollisionMask) Right() bool {
	return m.TopRight && m.BottomRight
}

func (m CollisionMask) count() (n int) {
	if m.TopLeft {
		n++
	}
	if m.TopRight {
		n++
	}
	if m.BottomLeft {
		n++
	}
	if m.BottomRight {
		n++
	}
	return n
}
