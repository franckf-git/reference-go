package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	const chessboardSquaresPerSide = 8
	const chessboardSquareSideSize = 120

	chessboardBlackSquareRects := make([]image.Rectangle, 0)
	chessboardWhiteSquareRects := make([]image.Rectangle, 0)

	for x := 0; x < chessboardSquaresPerSide; x++ {
		for y := 0; y < chessboardSquaresPerSide; y++ {
			x0 := x * chessboardSquareSideSize
			y0 := y * chessboardSquareSideSize
			x1 := x0 + chessboardSquareSideSize
			y1 := y0 + chessboardSquareSideSize

			rect := image.Rect(x0, y0, x1, y1)

			if isBlack := (x+y)%2 == 1; isBlack {
				chessboardBlackSquareRects = append(chessboardBlackSquareRects, rect)
			} else {
				chessboardWhiteSquareRects = append(chessboardWhiteSquareRects, rect)
			}
		}
	}

	const chessboardSideSize = chessboardSquaresPerSide * chessboardSquareSideSize

	chessboardRect := image.Rect(0, 0, chessboardSideSize, chessboardSideSize)
	chessboardImage := image.NewRGBA(chessboardRect)

	blackUniform := image.Uniform{color.Black}
	whiteUniform := image.Uniform{color.White}

	for _, rect := range chessboardBlackSquareRects {
		draw.Draw(chessboardImage, rect, &blackUniform, image.Point{}, draw.Src)
	}

	for _, rect := range chessboardWhiteSquareRects {
		draw.Draw(chessboardImage, rect, &whiteUniform, image.Point{}, draw.Src)
	}

	chessboardFile, err := os.Create("chessboard.png")
	if err != nil {
		panic(err)
	}

	png.Encode(chessboardFile, chessboardImage)
}
