package kata

func Solve(s string) []int {
	r := []rune(s)
	uppercase := 0
	lowercase := 0
	numbers := 0
	special := 0
	for _, v := range r {

		if 65 <= v && v <= 90 {
			uppercase++
		} else if 97 <= v && v <= 122 {
			lowercase++
		} else if 48 <= v && v <= 58 {
			numbers++
		} else {
			special++
		}

	}
	result := []int{uppercase, lowercase, numbers, special}
	return result
}

// top

package kata
func Solve(s string) []int {
  r := make([]int, 4)
  for _, c := range s {
    switch {
      case c >= 'A' && c <= 'Z': r[0]++;
      case c >= 'a' && c <= 'z': r[1]++;
      case c >= '0' && c <= '9': r[2]++;
      default: r[3]++;
    }
  }
  return r
}

// ---

package kata
import (
    "unicode"
  )

func Solve(s string) []int {
  var results []int = []int{0,0,0,0}

  for _, char := range s { //

    if(unicode.IsUpper(char)) {
      results[0]++
    } else if (unicode.IsLower(char)) {
      results[1]++
    } else if (unicode.IsDigit(char)) {
      results[2]++
    } else {
      results[3]++
    }
  }
  return results
}
