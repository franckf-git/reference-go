package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	r int = iota
	g
	b
)

func main() {
	var datas [100]int
	for i := range datas {
		datas[i] = rand.Intn(100)
	}
	datasToPng(datas)
}

func transform(nums [100]int) (colrs [100]color.RGBA) {
	var biggest int
	for _, v := range nums {
		if v >= biggest {
			biggest = v
		}
	}

	var curColrs color.RGBA
	for i, v := range nums {
		cur := (float64(v) / float64(biggest)) * 255
		rng := rand.Intn(3)
		switch rng {
		case r:
			curColrs = color.RGBA{
				R: uint8(cur),
				G: uint8(255),
				B: uint8(255),
				A: 255,
			}
		case g:
			curColrs = color.RGBA{
				R: uint8(255),
				G: uint8(cur),
				B: uint8(255),
				A: 255,
			}
		case b:
			curColrs = color.RGBA{
				R: uint8(255),
				G: uint8(255),
				B: uint8(cur),
				A: 255,
			}
		}
		colrs[i] = curColrs
	}
	return
}

func datasToPng(datas [100]int) {
	workingImage := image.NewRGBA(image.Rect(0, 0, 1000, 1000))
	datasMatrixed := matrix(datas)

	for line := range datasMatrixed {
		for col := range datasMatrixed[line] {
			rect := image.Rectangle{
				Min: image.Point{
					X: line * len(datas),
					Y: col * len(datas),
				},
				Max: image.Point{
					X: len(datas) + (line * len(datas)),
					Y: len(datas) + (col * len(datas)),
				},
			}

			coloring := image.Uniform{datasMatrixed[line][col]}

			draw.Draw(workingImage, rect, &coloring, image.Point{}, draw.Src)
		}
	}

	exportPNG, err := os.Create("export.png")
	if err != nil {
		panic(err)
	}
	png.Encode(exportPNG, workingImage)
}

func matrix(datas [100]int) (matrix [10][10]color.RGBA) {
	datasColored := transform(datas)
	var delta int
	for line := range matrix {
		for col := range matrix[line] {
			matrix[line][col] = datasColored[delta]
			delta++
		}
	}
	return
}

func valueToRGB(number int) (rgb color.RGBA) {
	avg := 256%(number+1) + number
	rgb = color.RGBA{
		R: uint8(avg),
		G: uint8(avg),
		B: uint8(avg),
		A: 255,
	}
	return
}
