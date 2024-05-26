package scene

import "github.com/nanore07/raytracer/types"

type ChequeredMaterial struct {
	// material has color and properties which tells how it reacts to light
	Color1          *types.Vec3
	Color2          *types.Vec3
	Ambient         float32
	Diffuse         float32
	Specular        float32
	ReflectionCoeff float32
}

func NewChequeredMaterial(
	color1 *types.Vec3, color2 *types.Vec3, ambient, diffuse, specular float32, reflectionCoeff float32,
) *ChequeredMaterial {
	return &ChequeredMaterial{
		Color1:          color1,
		Color2:          color2,
		Ambient:         ambient,
		Specular:        specular,
		Diffuse:         diffuse,
		ReflectionCoeff: reflectionCoeff,
	}
}

func (m *ChequeredMaterial) ColorAt(position types.Vec3) *types.Vec3 {
	if int((position.X+5.0)*3.0)%2 == int((position.Z+5.0)*3.0)%2 {
		return m.Color1
	}
	return m.Color2
}

func (m *ChequeredMaterial) GetDiffuse() float32 {
	return m.Diffuse
}

func (m *ChequeredMaterial) GetSpecular() float32 {
	return m.Specular
}

func (m *ChequeredMaterial) GetReflectionCoeff() float32 {
	return m.ReflectionCoeff
}

func (m *ChequeredMaterial) GetAmbient() float32 {
	return m.Ambient
}
