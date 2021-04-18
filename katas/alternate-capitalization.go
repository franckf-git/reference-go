package kata

import "strings"

func Capitalize(st string) []string {
	result := make([]string, 2)

	for i, _ := range st {
		if i%2 == 0 {
			//even
			result[0] = result[0] + strings.ToUpper(string(st[i]))
			result[1] = result[1] + string(st[i])
		} else {
			//odd
			result[0] = result[0] + string(st[i])
			result[1] = result[1] + strings.ToUpper(string(st[i]))
		}
	}

	return result
}

// top

package kata
import "unicode"
func Capitalize(s string) []string {
  a, b := []rune(s),[]rune(s)
  for i := range a {
    if i%2 == 0 {
      a[i] = unicode.ToUpper(a[i])
    }else{
      b[i] = unicode.ToUpper(b[i])
    }
  }
  return []string{string(a), string(b)}
}
