package types

import (
	"image/color"
	"log"
	"math"
	"strconv"
)

type Vec3 struct {
	X float32
	Y float32
	Z float32
}

func (v Vec3) Add(u Vec3) Vec3 {
	return Vec3{X: v.X + u.X, Y: v.Y + u.Y, Z: v.Z + u.Z}
}

func (v Vec3) Subtract(u Vec3) Vec3 {
	return Vec3{X: v.X - u.X, Y: v.Y - u.Y, Z: v.Z - u.Z}
}

func (v Vec3) Multiply(t float32) Vec3 {
	return Vec3{X: v.X * t, Y: v.Y * t, Z: v.Z * t}
}

func (v Vec3) Divide(t float32) Vec3 {
	return Vec3{X: v.X / t, Y: v.Y / t, Z: v.Z / t}
}

func (v Vec3) Dot(u Vec3) float32 {
	return v.X*u.X + v.Y*u.Y + v.Z*u.Z
}

func (v Vec3) Cross(u Vec3) Vec3 {
	return Vec3{
		X: v.Y*u.Z - v.Z*u.Y,
		Y: v.Z*u.X - v.X*u.Z,
		Z: v.X*u.Y - v.Y*u.X,
	}
}

func (v Vec3) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3) Normalize() Vec3 {
	len := v.Length()
	return Vec3{X: v.X / len, Y: v.Y / len, Z: v.Z / len}
}

func (v Vec3) ToColor() color.RGBA {
	return color.RGBA{
		R: uint8(math.Round(math.Max(math.Min(float64(v.X)*255, 255), 0))),
		G: uint8(math.Round(math.Max(math.Min(float64(v.Y)*255, 255), 0))),
		B: uint8(math.Round(math.Max(math.Min(float64(v.Z)*255, 255), 0))),
		A: 255,
	}
}

func convertToFloat(hexVal string) float32 {
	hexInt, err := strconv.ParseInt(hexVal, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return float32(hexInt) / 255.0
}

func NewVec3FromHex(hexCode string) *Vec3 {
	return &Vec3{X: convertToFloat(hexCode[1:3]), Y: convertToFloat(hexCode[3:5]), Z: convertToFloat(hexCode[5:7])}
}
