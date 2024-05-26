package scene

import (
	"math"

	"github.com/nanore07/raytracer/types"
)

type Sphere struct {
	Center   *types.Vec3
	Radius   float32
	Material Material
}

func NewSphere(Center *types.Vec3, Radius float32, Material Material) *Sphere {
	return &Sphere{
		Center:   Center,
		Radius:   Radius,
		Material: Material,
	}
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

func (s *Sphere) Normal(surface_point types.Vec3) types.Vec3 {
	// returns surface normal to the point on sphere's surface
	return surface_point.Subtract(*s.Center).Normalize()
}

func (s *Sphere) GetMaterial() Material {
	return s.Material
}
