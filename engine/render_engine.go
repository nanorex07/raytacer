package engine

import (
	"image/color"
	"math"

	"github.com/nanore07/raytracer/scene"
	"github.com/nanore07/raytracer/types"
)

type RenderEngine struct {
	// render 3d objects into 2d objects
}

func NewRenderEngine() *RenderEngine {
	return &RenderEngine{}
}

func (r *RenderEngine) Render(scene *scene.Scene) *types.Frame {

	floatW := float32(scene.Width)
	floatH := float32(scene.Height)

	aspect_ratio := floatW / floatH

	x0 := float32(-1.0)
	x1 := float32(1.0)

	xstep := (x1 - x0) / (floatW - 1.0)

	y0 := -1.0 / aspect_ratio
	y1 := -1.0 * y0

	ystep := (y1 - y0) / (floatH - 1)

	pixels := types.NewFrame(int(scene.Width), int(scene.Height))

	pixels.EachPixel(func(i int, j int) color.RGBA {
		y := y0 + float32(j)*ystep
		x := x0 + float32(i)*xstep
		dir := (&types.Vec3{X: x, Y: y, Z: 0}).Subtract(*scene.Camera)
		ray := types.NewRay(scene.Camera, &dir)
		return r.RayTrace(ray, scene).ToColor()
	})
	// for j := 0; j < int(scene.Height); j++ {
	// 	y := y0 + float32(j)*ystep
	// 	for i := 0; i < int(scene.Width); i++ {
	// 		x := x0 + float32(i)*xstep
	// 		dir := (&types.Vec3{X: x, Y: y, Z: 0}).Subtract(*scene.Camera)
	// 		ray := types.NewRay(scene.Camera, &dir)
	// 		pixels.SetPixel(i, j, r.RayTrace(ray, scene).ToColor())
	// 	}
	// }
	return pixels
}

func (r *RenderEngine) RayTrace(ray *types.Ray, scn *scene.Scene) types.Vec3 {
	color := types.Vec3{X: 0, Y: 0, Z: 0}
	dist_hit, obj_hit := r.FindNearest(ray, scn)
	if obj_hit == nil {
		return color
	}
	hit_pos := ray.Origin.Add(ray.Direction.Multiply(dist_hit))
	color = color.Add(r.ColorAt(obj_hit, hit_pos, scn))
	return color
}

func (r *RenderEngine) FindNearest(ray *types.Ray, scn *scene.Scene) (float32, scene.Object) {
	distMin := float32(math.Inf(1))
	var obj_hit scene.Object
	for _, obj := range scn.Objects {
		dist, is_hit := obj.Intersects(ray)
		if is_hit && (obj_hit != nil || dist < distMin) {
			distMin = dist
			obj_hit = obj
		}
	}
	return distMin, obj_hit
}

func (r *RenderEngine) ColorAt(obj_hit scene.Object, hit_pos types.Vec3, scn *scene.Scene) types.Vec3 {
	return obj_hit.ColorAt(&hit_pos)
}
