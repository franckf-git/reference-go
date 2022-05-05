package main

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	size := img.Bounds().Size()
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)
			fmt.Printf(
				"RED: %d, GREEN: %d, BLUE: %d, ALPHA: %d\n",
				col.R, col.G, col.B, col.A,
			)
		}
	}
}
