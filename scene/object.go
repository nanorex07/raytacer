package scene

import "github.com/nanore07/raytracer/types"

type Object interface {
	ColorAt(*types.Vec3) types.Vec3
	Intersects(*types.Ray) (float32, bool)
}
