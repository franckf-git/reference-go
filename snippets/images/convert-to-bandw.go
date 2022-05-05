package main

func toBlackAndWhite(originalImage image.Image, whiteThreshold uint8) image.Image {
	originalImage = toGreyscale(originalImage)
	size := originalImage.Bounds().Size()
	rect := image.Rect(0, 0, size.X, size.Y)
	modifiedImg := image.NewRGBA(rect)
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := originalImage.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)
			modifiedColorValue := originalColor.R
			if modifiedColorValue >= whiteThreshold {
				modifiedColorValue = 255
			} else {
				modifiedColorValue = 0
			}
			modifiedColor := color.RGBA{
				R: modifiedColorValue,
				G: modifiedColorValue,
				B: modifiedColorValue,
				A: originalColor.A,
			}
			modifiedImg.Set(x, y, modifiedColor)
		}
	}
	return modifiedImg
}
