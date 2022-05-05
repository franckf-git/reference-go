package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

func toGreyscale(originalImage image.Image) image.Image {
	size := originalImage.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	modifiedImg := image.NewRGBA(rect)
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := originalImage.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			red := float64(originalColor.R)
			green := float64(originalColor.G)
			blue := float64(originalColor.B)
			grey := uint8(
				math.Round((red + green + blue) / 3),
			)
			modifiedColor := color.RGBA{
				R: grey,
				G: grey,
				B: grey,
				A: originalColor.A,
			}
			modifiedImg.Set(x, y, modifiedColor)
		}
	}
	return modifiedImg
}

func main() {
	inputFile, err := os.Open("input.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer inputFile.Close()
	img, _, err := image.Decode(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	greyImg := toGreyscale(img)
	outputFile, err := os.Create("output.jpg")
	if err != nil {
		log.Fatalln(err)
	}
	defer outputFile.Close()
	err = jpeg.Encode(outputFile, greyImg, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
