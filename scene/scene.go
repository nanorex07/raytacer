package scene

import "github.com/nanore07/raytracer/types"

type Scene struct {
	Camera  *types.Vec3
	Objects []Object
	Lights  []*Light
	Width   int32
	Height  int32
}

func NewScene(camera *types.Vec3, objects []Object, lights []*Light, width int32, height int32) *Scene {
	return &Scene{
		Camera:  camera,
		Objects: objects,
		Lights:  lights,
		Width:   width,
		Height:  height,
	}
}
