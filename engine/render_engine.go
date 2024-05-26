package engine

import (
	"image/color"
	"math"

	"github.com/nanore07/raytracer/scene"
	"github.com/nanore07/raytracer/types"
)

type RenderEngine struct {
	// render 3d objects into 2d objects
	MAX_DEPTH    int
	MIN_DISPLACE float32
}

func NewRenderEngine(max_depth int, min_displace float32) *RenderEngine {
	return &RenderEngine{
		MAX_DEPTH:    max_depth,
		MIN_DISPLACE: min_displace,
	}
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
		return r.RayTrace(ray, scene, r.MAX_DEPTH).ToColor()
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

func (r *RenderEngine) RayTrace(ray *types.Ray, scn *scene.Scene, depth int) types.Vec3 {
	color := *types.NewVec3FromHex("#111111")
	dist_hit, obj_hit := r.FindNearest(ray, scn)
	if obj_hit == nil {
		return color
	}
	hit_pos := ray.Origin.Add(ray.Direction.Multiply(dist_hit))
	hit_normal := obj_hit.Normal(hit_pos)
	color = color.Add(r.ColorAt(obj_hit, hit_pos, hit_normal, scn))
	if depth > 0 {
		new_ray_pos := hit_pos.Add(hit_normal.Multiply(r.MIN_DISPLACE))
		new_ray_dir := ray.Direction.Subtract(hit_normal.Multiply(2 * ray.Direction.Dot(hit_normal)))
		new_ray := types.NewRay(&new_ray_pos, &new_ray_dir)

		// Attenuate the reflected ray by reflection coefficient
		color = color.Add(r.RayTrace(new_ray, scn, depth-1)).Multiply(obj_hit.GetMaterial().GetReflectionCoeff())
	}
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

func (r *RenderEngine) ColorAt(obj_hit scene.Object, hit_pos types.Vec3, hit_normal types.Vec3, scn *scene.Scene) types.Vec3 {
	material := obj_hit.GetMaterial()
	obj_color := material.ColorAt(hit_pos)

	specular_k := 60

	to_cam := scn.Camera.Subtract(hit_pos)
	color := types.NewVec3FromHex("#FFFFFF").Multiply(material.GetAmbient())

	for _, light := range scn.Lights {
		to_light_dir := light.Position.Subtract(hit_pos)
		to_light := types.NewRay(&hit_pos, &to_light_dir)

		// Diffuse shading (Lambert)
		color = color.Add(obj_color.Multiply(
			material.GetDiffuse() * float32(
				math.Max(float64(hit_normal.Dot(*to_light.Direction)), 0),
			),
		))
		// Specular Shading (Blinn-Phong)
		half_vector := to_light.Direction.Add(to_cam).Normalize()
		color = color.Add(light.Color.Multiply(material.GetSpecular()).Multiply(
			float32(math.Pow(math.Max(float64(hit_normal.Dot(half_vector)), 0), float64(specular_k))),
		))
	}

	return color
}
