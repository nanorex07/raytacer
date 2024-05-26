package main

import (
	"github.com/nanore07/raytracer/engine"
	"github.com/nanore07/raytracer/scene"
	"github.com/nanore07/raytracer/types"
)

const (
	WIDTH  = 320
	HEIGHT = 200
)

func main() {

	camera := types.Vec3{X: 0, Y: 0, Z: -1}
	objects := []scene.Object{
		scene.NewSphere(&types.Vec3{X: 0, Y: 0, Z: 0}, 0.5, types.NewVec3FromHex("#0000FF")),
	}
	scene := scene.NewScene(&camera, objects, WIDTH, HEIGHT)
	engine := engine.NewRenderEngine()
	frame := engine.Render(scene)
	frame.Save("output.png")
}
