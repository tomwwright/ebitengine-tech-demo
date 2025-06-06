package factories

import (
	"fmt"

	"github.com/solarlune/resolv"
	"github.com/tomwwright/ebitengine-tech-demo/components"
	"github.com/tomwwright/ebitengine-tech-demo/components/collision"
	"github.com/tomwwright/ebitengine-tech-demo/constants"
	"github.com/tomwwright/ebitengine-tech-demo/tags"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/transform"
)

func AddCollision(entry *donburi.Entry, collision collision.CollisionType) error {
	objects := collision.Mask().Objects()
	if objects == nil || len(objects) == 0 {
		return fmt.Errorf("invalid collision type mapped to no objects: %s", collision)
	}

	// scale by tile size
	for i, o := range objects {
		objects[i].Position = o.Position.MulScalar(constants.TileSize)
		objects[i].Size = o.Size.MulScalar(constants.TileSize)
	}

	// if the first object has no offset it can be added directly to the entry
	first := objects[0]
	if first.Position == constants.Zero {
		object := resolv.NewObject(0, 0, first.Size.X, first.Size.Y, tags.ResolvTagCollider)
		object.Data = object
		entry.AddComponent(components.Object)
		components.Object.Set(entry, object)

		objects = objects[1:]
	}

	// create child entries for additional collision objects
	for _, o := range objects {
		e := entry.World.Create(components.Transform, components.Object)
		child := entry.World.Entry(e)

		t := components.Transform.Get(child)
		t.LocalPosition = o.Position

		object := resolv.NewObject(0, 0, o.Size.X, o.Size.Y, tags.ResolvTagCollider)
		object.Data = child
		components.Object.Set(child, object)

		transform.AppendChild(entry, child, false)
	}

	return nil
}
