package main

import (
	"fmt"
	"math/bits"
	"regexp"
	"strconv"
)

const TileSize = 10    // Size of a tile
const TileOpposite = 9 // "Opposite" border (i.e. right or bottom)

type Tile struct {
	ID      TileID
	Image   Image
	Borders [8]Border
}

type TileID int
type Border uint16

const (
	BorderTop       = iota // Top border (left to right)
	BorderRight            // Right border (top to bottom)
	BorderBottom           // Bottom border (left to right)
	BorderLeft             // Left border (top to bottom)
	BorderTopRev           // Reversed top border
	BorderRightRev         // Reversed right border
	BorderBottomRev        // Reversed bottom border
	BorderLeftRev          // Reversed left border
)

var tileHeaderFmt = regexp.MustCompile(`^Tile (\d{4}):$`)

func NewTileFromData(data []string) *Tile {
	if len(data) != TileSize+1 {
		panic(fmt.Errorf("invalid tile size: %d", len(data)))
	}
	match := tileHeaderFmt.FindStringSubmatch(data[0])
	if match == nil {
		panic(fmt.Errorf("invalid tile header: %q", data[0]))
	}
	id, _ := strconv.Atoi(match[1])
	t := Tile{
		ID:    TileID(id),
		Image: ImageFromData(data[1:]),
	}
	t.computeBorders()
	return &t
}

func (t *Tile) computeBorders() {
	// Clear borders before recomputing them
	for i := 0; i < len(t.Borders); i++ {
		t.Borders[i] = 0
	}
	for x := 0; x < TileSize; x++ {
		bit := Border(1) << x
		if t.Image[0][x] == '#' {
			t.Borders[BorderTop] |= bit
		}
		if t.Image[x][TileOpposite] == '#' {
			t.Borders[BorderRight] |= bit
		}
		if t.Image[TileOpposite][x] == '#' {
			t.Borders[BorderBottom] |= bit
		}
		if t.Image[x][0] == '#' {
			t.Borders[BorderLeft] |= bit
		}
	}
	t.Borders[BorderTopRev] = reverseBorder(t.Borders[BorderTop])
	t.Borders[BorderRightRev] = reverseBorder(t.Borders[BorderRight])
	t.Borders[BorderBottomRev] = reverseBorder(t.Borders[BorderBottom])
	t.Borders[BorderLeftRev] = reverseBorder(t.Borders[BorderLeft])
}

func reverseBorder(b Border) Border {
	return Border(bits.Reverse16(uint16(b)) >> (16 - TileSize))
}

func (t *Tile) findBorder(b Border) int {
	for i, border := range t.Borders {
		if b == border {
			return i
		}
	}
	panic("Border not found")
}

// ToLeftBorder transforms the tile so that its left border
// aligns with given border
func (t *Tile) ToLeftBorder(b Border) {
	switch t.findBorder(b) {
	case BorderLeft:
		return
	case BorderLeftRev:
		t.Image.FlipV()
	case BorderTop:
		t.Image.FlipD()
	case BorderTopRev:
		t.Image.RotateLeft()
	case BorderRight:
		t.Image.FlipH()
	case BorderRightRev:
		t.Image.FlipH().FlipV()
	case BorderBottom:
		t.Image.RotateRight()
	case BorderBottomRev:
		t.Image.RotateRight().FlipV()
	}
	t.computeBorders()
}

// ToTopBorder transforms the tile so that its top border aligns
// with given border.
func (t *Tile) ToTopBorder(b Border) {
	switch t.findBorder(b) {
	case BorderTop:
		return
	case BorderTopRev:
		t.Image.FlipH()
	case BorderRight:
		t.Image.RotateLeft()
	case BorderRightRev:
		t.Image.RotateLeft().FlipH()
	case BorderBottom:
		t.Image.FlipV()
	case BorderBottomRev:
		t.Image.FlipV().FlipH()
	case BorderLeft:
		t.Image.FlipD()
	case BorderLeftRev:
		t.Image.RotateRight()
	}
	t.computeBorders()
}

// ToTopLeftCorner transforms the tile until it fits the top left corner
func (t *Tile) ToTopLeftCorner(index BorderIndex) {
	top, bottom, left, right := t.findUniqueBorders(index)
	switch true {
	case top && left:
		return
	case top && right:
		t.Image.FlipH()
	case bottom && left:
		t.Image.FlipV()
	case bottom && right:
		t.Image.FlipV().FlipH()
	default:
		// Unreachable unless t isn't a corner
		panic("Tile isn't a corner")
	}
	t.computeBorders()
}

func (t *Tile) findUniqueBorders(index BorderIndex) (top, bottom, left, right bool) {
	// Lookup only the 4 first borders
	for idx, b := range t.Borders[:4] {
		if len(index[b]) > 1 {
			continue
		}
		switch idx {
		case BorderTop:
			top = true
		case BorderBottom:
			bottom = true
		case BorderLeft:
			left = true
		case BorderRight:
			right = true
		}
	}
	return
}
