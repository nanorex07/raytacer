package types

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
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
	wg := sync.WaitGroup{}
	goroutineSemaphore := make(chan struct{}, 10000)

	for x := 0; x < s.Width; x++ {
		for y := 0; y < s.Height; y++ {
			wg.Add(1)
			goroutineSemaphore <- struct{}{}
			go func(x, y int) {
				defer wg.Done()
				defer func() { <-goroutineSemaphore }()
				s.Img.Set(x, y, colorFunction(x, y))
			}(x, y)
		}
	}
	wg.Wait()
}
