package main

import (
	"fmt"
	"math"
)

// Tiles are indexed by their ID.
type TileSet map[TileID]*Tile

// BorderIndex maps the (1024) border patterns to the matching tiles.
type BorderIndex [][]TileID

func parseTiles(input []string) TileSet {
	tiles := make(TileSet)
	for i := 0; i < len(input); i++ {
		line := input[i]
		if line == "" {
			continue
		}
		t := NewTileFromData(input[i : i+TileSize+1])
		i += TileSize + 1
		tiles[t.ID] = t
	}
	return tiles
}

func makeBorderIndex(tiles TileSet) BorderIndex {
	index := make(BorderIndex, 1<<TileSize)
	for id, tile := range tiles {
		for _, b := range tile.Borders {
			index[b] = append(index[b], id)
		}
	}
	return index
}

// Reorder tiles, so they can be assembled to an image.
// Returns the width (in number of tiles) of the image and the ordered tiles.
func reorder(tiles TileSet, index BorderIndex, topLeftID TileID) (int, []*Tile) {
	result := make([]*Tile, len(tiles))
	width := int(math.Sqrt(float64(len(tiles))))
	if width*width != len(tiles) {
		panic(fmt.Errorf("invalid width: %d", width))
	}
	// Initialize with the top left corner
	topLeft := tiles[topLeftID]
	topLeft.ToTopLeftCorner(index)
	result[0] = topLeft
	for i := 1; i < len(result); i++ {
		var border Border
		var prev, current TileID
		if i%width == 0 {
			// Beginning of a row, match the bottom border of the tile above
			prev = result[i-width].ID
			border = result[i-width].Borders[BorderBottom]
		} else {
			// Match the right border of the previous tile
			prev = result[i-1].ID
			border = result[i-1].Borders[BorderRight]
		}
		matching := index[border]
		if len(matching) != 2 {
			panic(fmt.Errorf("%d tiles with matching border", len(matching)))
		}
		if matching[0] == prev {
			current = matching[1]
		} else {
			current = matching[0]
		}

		t := tiles[current]
		// Transform tile to put it in place
		if i%width == 0 {
			t.ToTopBorder(border)
		} else {
			t.ToLeftBorder(border)
		}
		result[i] = t
	}
	return width, result
}

// Assemble an image from the ordered tiles
func assemble(tiles TileSet, index BorderIndex, topLeftID TileID) Image {
	width, orderedTiles := reorder(tiles, index, topLeftID)
	blitWidth := orderedTiles[0].Image.Width() - 2
	imgWidth := width * blitWidth
	img := MakeImage(imgWidth, imgWidth)

	var blitX, blitY int
	for _, tile := range orderedTiles {
		if blitX == img.Width() {
			blitX = 0
			blitY += blitWidth
		}
		for y := 0; y < tile.Image.Height()-2; y++ {
			copy(img[blitY+y][blitX:blitX+blitWidth], tile.Image[y+1][1:blitWidth+1])
		}
		blitX += blitWidth
	}
	return img
}
