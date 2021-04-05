package kata

import "strconv"

func countSheep(num int) string {
  var result string
  if num == 0 {
    return ""
  }
	for i := 1; i <= num; i++ {
		result = result + strconv.Itoa(i) + " sheep..."
	}
	return result
}

// top

package kata

import (
  "fmt"
  "strings"
)

func countSheep(num int) string {
  var sb strings.Builder

  for count := 1; count <= num; count++ {
        fmt.Fprintf(&sb, "%d sheep...", count)
  }

  return sb.String()
}
