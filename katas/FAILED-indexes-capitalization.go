// mine

package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string, arr []int) string {
	var result string = st
	for num := 0; num < len(arr); num++ {
		for i, v := range st {
			if i == arr[num] {
				result = result + strings.ToUpper(string(v))
			} else {
				result = result + string(v)
			}
		}

	}
	return result
}

func main() {

	fmt.Print(Capitalize("abcdef", []int{1, 2, 5})) //abcdefaBcdefabCdefabcdeF
}


// top

package kata
import("unicode")
func Capitalize(s string, a []int) string {
  r := []rune(s)
  for _, v := range a {
    if v>=0 && v<len(r) {
      r[v] = unicode.ToUpper(r[v])
    }
  }
  return string(r)
}

// ---

package kata

func Capitalize(s string, arr []int) string {
  c := []rune(s)
  for _, i := range arr {
    if i < len(s) {c[i] = rune(s[i] - 32)}
  }
  return string(c)
}
