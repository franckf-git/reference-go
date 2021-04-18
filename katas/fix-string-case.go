package kata

import "strings"

func solve(str string) string {
	data := []rune(str)
	var uppercas int
	var lowercas int
	for _, v := range data {
		if v > 96 {
			lowercas++
		} else {
			uppercas++
		}
	}

	if lowercas == 0 || uppercas > lowercas {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}

}

// top

package kata

import (
  "unicode"
  "strings"
)

func solve(str string) string {
  uppers := 0

  for _, rune := range str {
    if (unicode.IsUpper(rune)) {
      uppers += 1
    }
  }

  if (uppers > len(str) / 2) {
    return strings.ToUpper(str)
  }

  return strings.ToLower(str)
}

// ---

package kata
import "strings"

func solve(str string) (result string) {
  uppercase, lowercase := 0, 0

  for _, v := range str {
    if v >= 65 && v <= 90 {
      uppercase++
    } else {
      lowercase++
    }
  }
  if uppercase > lowercase {
    result = strings.ToUpper(str)
  } else {
    result = strings.ToLower(str)
  }
  return result
}
