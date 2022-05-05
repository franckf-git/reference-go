package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"sync"
)

func toGreyscale(originalImage image.Image) image.Image
func toBlackAndWhite(originalImage image.Image, whiteThreshold uint8) image.Image

func toBlackAndWhiteWithCalculatedWhiteThreshold(originalImage image.Image) image.Image {
	size := originalImage.Bounds().Size()

	totalPixels := size.X * size.Y
	halfPixels := int(math.Round(float64(totalPixels) / 2))

	var closestThreshold uint8
	var closestHalfPixelsDelta int
	var closestImage image.Image

	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 256; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			modifiedImg := toBlackAndWhite(originalImage, uint8(i))
			var blackPixels int

			for x := 0; x < size.X; x++ {
				for y := 0; y < size.Y; y++ {
					pixel := modifiedImg.At(x, y)
					col := color.RGBAModel.Convert(pixel).(color.RGBA)

					if col.R == 0 {
						blackPixels++
					}
				}
			}

			defer mutex.Unlock()
			mutex.Lock()

			halfPixelsDelta := int(math.Abs(float64(halfPixels - blackPixels)))

			if closestHalfPixelsDelta == 0 || halfPixelsDelta < closestHalfPixelsDelta {
				closestHalfPixelsDelta = halfPixelsDelta
				closestThreshold = uint8(i)
				closestImage = modifiedImg
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("CHOSEN THRESHOLD: %d\n", closestThreshold)

	return closestImage
}
