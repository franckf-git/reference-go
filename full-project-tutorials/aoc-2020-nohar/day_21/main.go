package main

import (
	"fmt"
	"math/bits"
	"sort"
	"strings"

	"gitlab.com/neuware/aoc-2020/utils"
)

type Allergen struct {
	Name       string // Name of the allergen
	Mask       uint   // Unique allergen mask
	Count      int    // Number of times this allergen appears
	Ingredient string // Ingredient containing this allergen
}

// ByName implements sort.Interface to sort allergens by name
type ByName []Allergen

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type OccurrenceTable map[string][]uint
type Candidates map[string]uint

func main() {
	table, allergens := parseInput(utils.ReadLines())

	candidates, part1 := analyze(table, allergens)
	fmt.Println("Part 1:", part1)

	refine(candidates, allergens)
	sort.Sort(ByName(allergens))
	part2 := make([]string, len(allergens))
	for i, a := range allergens {
		part2[i] = a.Ingredient
	}
	fmt.Println("Part 2:", strings.Join(part2, ","))
}

// Refine candidates until we obtain a 1-1 mapping between ingredients and the
// allergens they contain.
func refine(candidates Candidates, allergens []Allergen) {
	var known uint
	for i := 0; i < len(allergens); i++ {
		for ingredient, possibleAllergens := range candidates {
			possibleAllergens &^= known
			if bits.OnesCount(possibleAllergens) == 1 {
				// Found the allergen, set it
				known |= possibleAllergens
				set(allergens, possibleAllergens, ingredient)
			}
		}
	}
}

// Associate an ingredient to the corresponding allergen given its mask
func set(allergens []Allergen, mask uint, ingredient string) {
	for a, allergen := range allergens {
		if allergen.Mask == mask {
			allergens[a].Ingredient = ingredient
		}
	}
}

// Analyse an occurrence table to filter out ingredients that contain no
// allergen. Return a candidate table and the number of occurrences of
// non-allergens.
func analyze(table OccurrenceTable, allergens []Allergen) (Candidates, int) {
	var result int
	candidates := make(Candidates)
	for ingredient, masks := range table {
		// Compute a mask showing all possible allergens for this food
		var possibleAllergens uint
		for _, m := range masks {
			possibleAllergens |= m
		}

		for _, allergen := range allergens {
			if possibleAllergens&allergen.Mask == 0 {
				// This allergen isn't in any food that contains this
				// ingredient
				continue
			}
			// Count the co-occurrences of this allergen and this ingredient.
			var count int
			for _, mask := range masks {
				if allergen.Mask&mask != 0 {
					count++
				}
			}

			// If the count doesn't match the allergen's occurrences, then this
			// ingredient doesn't have it.
			if count != allergen.Count {
				possibleAllergens &^= allergen.Mask
			}
		}

		if possibleAllergens == 0 {
			result += len(masks)
		} else {
			candidates[ingredient] = possibleAllergens
		}
	}
	return candidates, result
}

func parseInput(input []string) (OccurrenceTable, []Allergen) {
	table := make(OccurrenceTable)
	allergens := make([]Allergen, 0, 8)
	// temporary index to find new allergens
	allergenIndex := make(map[string]int)
	for _, line := range input {
		s := strings.Split(line, " (contains ")
		if len(s) != 2 {
			panic("invalid input")
		}
		ingredients := strings.Fields(s[0])
		allergenList := strings.Split(strings.Trim(s[1], ") "), ", ")

		var mask uint
		for _, name := range allergenList {
			// Register new allergen if we haven't seen it before
			idx, ok := allergenIndex[name]
			if !ok {
				idx = len(allergens)
				allergens = append(allergens, Allergen{
					Name: name,
					Mask: 1 << idx,
				})
				allergenIndex[name] = idx
			}
			mask |= allergens[idx].Mask
			allergens[idx].Count++
		}
		for _, i := range ingredients {
			table[i] = append(table[i], mask)
		}
	}
	return table, allergens
}
