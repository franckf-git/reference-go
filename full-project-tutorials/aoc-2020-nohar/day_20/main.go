package main

import (
	"fmt"

	"gitlab.com/neuware/aoc-2020/utils"
)

var monsterPattern = ImageFromData([]string{
	"                  # ",
	"#    ##    ##    ###",
	" #  #  #  #  #  #   ",
})

func main() {
	tiles := parseTiles(utils.ReadLines())
	index := makeBorderIndex(tiles)
	corners := findCorners(tiles, index)
	fmt.Println("Corners:", corners)

	part1 := 1
	for _, id := range corners {
		part1 *= int(id)
	}
	fmt.Println("Part 1:", part1)

	img := assemble(tiles, index, corners[0])
	monsters := countMonsters(img)
	fmt.Println("Found", monsters, "monsters!")
	part2 := img.Count('#') - monsters*monsterPattern.Count('#')
	fmt.Println("Part 2:", part2)
}

func countMonsters(img Image) int {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if count := img.Find(monsterPattern); count != 0 {
				return count
			}
			img.FlipV()
			if count := img.Find(monsterPattern); count != 0 {
				return count
			}
			img.FlipH()
		}
		img.FlipD()
	}
	panic("pattern not found!")
}

func findCorners(tiles TileSet, index BorderIndex) []TileID {
	corners := make([]TileID, 0, 4)
	for _, tile := range tiles {
		// Count this tile's unique corners
		var unique int
		for _, b := range tile.Borders {
			if len(index[b]) == 1 {
				unique++
			}
		}
		// Corners have exactly 4 unique borders (counting the reversed ones)
		if unique == 4 {
			corners = append(corners, tile.ID)
		}
	}
	if len(corners) != 4 {
		panic(fmt.Errorf("Found %d corners!", len(corners)))
	}
	return corners
}
