package scene

import (
	"math"

	"github.com/nanore07/raytracer/types"
)

type Sphere struct {
	Center *types.Vec3
	Radius float32
	Color  *types.Vec3
}

func NewSphere(Center *types.Vec3, Radius float32, Color *types.Vec3) *Sphere {
	return &Sphere{
		Center: Center,
		Radius: Radius,
		Color:  Color,
	}
}

func (s *Sphere) ColorAt(coor *types.Vec3) types.Vec3 {
	return *s.Color
}

func (s *Sphere) Intersects(ray *types.Ray) (float32, bool) {
	// Checks if a ray intersects sphere returns distance to intersection
	// and bool for intersection
	sphere_to_ray := ray.Origin.Subtract(*s.Center)
	b := 2.0 * ray.Direction.Dot(sphere_to_ray)
	c := sphere_to_ray.Dot(sphere_to_ray) - s.Radius*s.Radius

	discriminant := b*b - 4*c
	if discriminant >= 0 {
		dist := (-b - float32(math.Sqrt(float64(discriminant)))) / 2.0
		if dist > 0 {
			return dist, true
		}
	}
	return 0, false
}
