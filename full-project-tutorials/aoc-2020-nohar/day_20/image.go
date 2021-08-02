package main

import "bytes"

// An Image holds pixel data
type Image [][]byte

// MakeImage initializes an empty image with given width and height
func MakeImage(width, height int) Image {
	img := make(Image, height)
	for y := 0; y < height; y++ {
		img[y] = make([]byte, width)
	}
	return img
}

// ImageFromData initializes an image from string data
func ImageFromData(data []string) Image {
	img := make(Image, len(data))
	for y, row := range data {
		img[y] = []byte(row)
	}
	return img
}

func (i Image) Width() int {
	return len(i[0])
}

func (i Image) Height() int {
	return len(i)
}

// Count returns the number of occurrences of given byte in the image
func (i Image) Count(p byte) int {
	var count int
	needle := []byte{p}
	for _, row := range i {
		count += bytes.Count(row, needle)
	}
	return count
}

// Find returns the number of occurrences of given pattern in the image
// Only '#' bytes are counted.
func (i Image) Find(pattern Image) int {
	var count int
	for y := 0; y < i.Height()-pattern.Height(); y++ {
	SearchPattern:
		for x := 0; x < i.Width()-pattern.Width(); x++ {
			for py := 0; py < pattern.Height(); py++ {
				for px := 0; px < pattern.Width(); px++ {
					if pattern[py][px] != '#' {
						continue
					}
					if i[y+py][x+px] != '#' {
						continue SearchPattern
					}
				}
			}
			// If we reach here, then the pattern was found
			count++
		}
	}
	return count
}

// Flip image vertically
func (i Image) FlipV() Image {
	height := len(i)
	for y := 0; y < height/2; y++ {
		i[y], i[height-y-1] = i[height-y-1], i[y]
	}
	return i
}

// Flip image horizontally
func (i Image) FlipH() Image {
	height, width := len(i), len(i[0])
	for y := 0; y < height; y++ {
		for x := 0; x < width/2; x++ {
			i[y][x], i[y][width-x-1] = i[y][width-x-1], i[y][x]
		}
	}
	return i
}

// Flip image diagonally (x <=> y) (assuming it's square)
func (i Image) FlipD() Image {
	height, width := len(i), len(i[0])
	if height != width {
		panic("image isn't square!")
	}
	for y := 0; y < height; y++ {
		for x := y; x < width; x++ {
			i[y][x], i[x][y] = i[x][y], i[y][x]
		}
	}
	return i
}

// Rotate image 90° to the right
func (i Image) RotateRight() Image {
	return i.FlipD().FlipH()
}

// Rotate image 90° to the left
func (i Image) RotateLeft() Image {
	return i.FlipD().FlipV()
}
