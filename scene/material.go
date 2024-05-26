package scene

import "github.com/nanore07/raytracer/types"

type Material interface {
	ColorAt(types.Vec3) *types.Vec3
	GetDiffuse() float32
	GetSpecular() float32
	GetAmbient() float32
	GetReflectionCoeff() float32
}

type PlainMaterial struct {
	// material has color and properties which tells how it reacts to light
	Color           *types.Vec3
	Ambient         float32
	Diffuse         float32
	Specular        float32
	ReflectionCoeff float32
}

func NewPlainMaterial(color *types.Vec3, ambient, diffuse, specular float32, reflectionCoeff float32) *PlainMaterial {
	return &PlainMaterial{
		Color:           color,
		Ambient:         ambient,
		Specular:        specular,
		Diffuse:         diffuse,
		ReflectionCoeff: reflectionCoeff,
	}
}

func (m *PlainMaterial) ColorAt(position types.Vec3) *types.Vec3 {
	return m.Color
}

func (m *PlainMaterial) GetDiffuse() float32 {
	return m.Diffuse
}

func (m *PlainMaterial) GetSpecular() float32 {
	return m.Specular
}

func (m *PlainMaterial) GetReflectionCoeff() float32 {
	return m.ReflectionCoeff
}

func (m *PlainMaterial) GetAmbient() float32 {
	return m.Ambient
}
