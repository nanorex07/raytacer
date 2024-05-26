package scene

import "github.com/nanore07/raytracer/types"

type Object interface {
	Intersects(*types.Ray) (float32, bool)
	Normal(types.Vec3) types.Vec3
	GetMaterial() Material
}
