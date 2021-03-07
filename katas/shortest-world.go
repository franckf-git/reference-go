package kata
import "strings"
func FindShort(s string) int {
  list := strings.Split(s, " ")
  var ref int
  for key,value := range list {
    if key == 0 {
      ref = len(value)
    }
    if len(value) < ref {
       ref = len(value)
    }
  }
  return ref
}

// top

package kata

import "strings"

func FindShort(s string) int {
  shortest := len(s)
  for _, word := range strings.Split(s, " ") {
    if len(word) < shortest {
      shortest = len(word)
    }
  }
  return shortest
}
