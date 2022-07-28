package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {

	canvas := image.NewRGBA(image.Rect(0, 0, 800, 800))
	draw.Draw(canvas, canvas.Bounds(), &image.Uniform{color.White}, image.Point{}, draw.Src)

	var grid [8][8]int
	for line := range grid {
		for col := range grid[line] {
			if line%2 == 0 {
				if col%2 != 0 {
					square := image.Rectangle{
						Min: image.Point{
							X: 100 * col,
							Y: 100 * line,
						},
						Max: image.Point{
							X: 100*col + 100,
							Y: 100*line + 100,
						},
					}
					draw.Draw(canvas, square, &image.Uniform{color.Black}, image.Point{}, draw.Src)
				}
			} else {
				if col%2 == 0 {
					square := image.Rectangle{
						Min: image.Point{
							X: 100 * col,
							Y: 100 * line,
						},
						Max: image.Point{
							X: 100*col + 100,
							Y: 100*line + 100,
						},
					}
					draw.Draw(canvas, square, &image.Uniform{color.Black}, image.Point{}, draw.Src)
				}
			}
		}
	}

	exportPNG, _ := os.Create("chessboard.png")
	png.Encode(exportPNG, canvas)
}
