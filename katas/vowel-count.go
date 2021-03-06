// tentative avec map

package kata

func GetCount(str string) (count int) {
	var vowels = map[string]int{
		"a": 0,
		"e": 0,
		"i": 0,
		"o": 0,
		"u": 0,
	}

	for _, value := range str {
		if vowels[string(value)] == 0 {
			vowels[string(value)] = vowels[string(value)] + 1
		}
	}
	//count := vowels["a"]

	return count
}

// solus

package kata

func GetCount(str string) (count int) {
	for _, value := range str {
		if string(value) == "a" || string(value) == "e" || string(value) == "i" || string(value) == "o" || string(value) == "u" {
			count++
		}
	}
	return count
}

// top

package kata

func GetCount(str string) (count int) {
  for _, c := range str {
    switch c {
    case 'a', 'e', 'i', 'o', 'u':
      count++
    }
  }
  return count
}

// ---

package kata

import "strings"

func GetCount(strn string) int {
  count := 0

  vowels := []string{"a", "e", "i", "o", "u"}

  for _, vowel := range vowels {
    count += strings.Count(strn, vowel)
  }

  return count
}