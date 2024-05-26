package main

import (
	"github.com/nanore07/raytracer/engine"
	"github.com/nanore07/raytracer/scene"
	"github.com/nanore07/raytracer/types"
)

const (
	WIDTH  = 3200
	HEIGHT = 4300
)

func main() {

	camera := types.Vec3{X: 0, Y: -0.35, Z: -1}
	objects := []scene.Object{

		scene.NewSphere(
			&types.Vec3{X: 0, Y: 10000.5, Z: 1}, 10000.0,
			scene.NewChequeredMaterial(types.NewVec3FromHex("#420500"), types.NewVec3FromHex("#E6B87D"), 0.05, 1.0, 1.0, 0.2),
		),

		scene.NewSphere(
			&types.Vec3{X: 0.75, Y: -0.1, Z: 1}, 0.6,
			scene.NewPlainMaterial(types.NewVec3FromHex("#AEFF11"), 0.05, 1.0, 1.0, 0.5),
		),

		scene.NewSphere(
			&types.Vec3{X: -0.75, Y: -0.1, Z: 2.25}, 0.6,
			scene.NewPlainMaterial(types.NewVec3FromHex("#8043EA"), 0.05, 1.0, 1.0, 0.5),
		),

		scene.NewSphere(
			&types.Vec3{X: -1.20, Y: 0.1, Z: 1.0}, 0.3,
			scene.NewPlainMaterial(types.NewVec3FromHex("#FF0000"), 0.05, 1.0, 1.0, 0.5),
		),
	}
	lights := []*scene.Light{
		scene.NewLight(&types.Vec3{X: 1, Y: -0.5, Z: -10.0}, types.NewVec3FromHex("#FFFFFF")),
		scene.NewLight(&types.Vec3{X: -0.5, Y: -10.5, Z: 0.0}, types.NewVec3FromHex("#FFFFFF")),
	}

	scene := scene.NewScene(&camera, objects, lights, WIDTH, HEIGHT)
	engine := engine.NewRenderEngine(
		6, 0.0001,
	)
	frame := engine.Render(scene)
	frame.Save("output.png")
}
