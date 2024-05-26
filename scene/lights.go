package scene

import "github.com/nanore07/raytracer/types"

type Light struct {
	Position *types.Vec3
	Color    *types.Vec3
}

func NewLight(position, color *types.Vec3) *Light {
	return &Light{
		Position: position,
		Color:    color,
	}
}
