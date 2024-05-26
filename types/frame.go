package types

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type Frame struct {
	Width, Height int
	Img           *image.RGBA
}

func NewFrame(width int, height int) *Frame {
	return &Frame{
		Width:  width,
		Height: height,
		Img:    image.NewRGBA(image.Rect(0, 0, width, height)),
	}
}

func (s *Frame) Save(filename string) {
	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, s.Img)
}

func (s *Frame) SetPixel(x int, y int, color color.RGBA) {
	s.Img.Set(x, y, color)
}

func (s *Frame) EachPixel(colorFunction func(int, int) color.RGBA) {
	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			s.Img.Set(x, y, colorFunction(x, y))
		}
	}
}
