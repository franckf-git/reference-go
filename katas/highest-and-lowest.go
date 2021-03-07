package kata
import (
  "strings"
  "strconv"
  "sort"
)

func HighAndLow(in string) string {
  arrayin := strings.Split(in, " ")
  inint := []int{}
  for _, value := range arrayin {
    valueint,_ := strconv.Atoi(value)
    newEntry := []int{valueint}
    inint = append(inint, newEntry...)
  }
  sort.Ints(inint)
  top := strconv.Itoa(inint[0])
  last := strconv.Itoa(inint[len(inint)-1])
  return last + " " + top
}

// top

package kata

import (
  "strings"
  "strconv"
  "fmt"
)

func HighAndLow(in string) string {
  var tmpH, tmpL int
  for i, s := range strings.Fields(in) {
    n, _ := strconv.Atoi(string(s))
    if i == 0 {
      tmpH = n
      tmpL = n
    }

    if n > tmpH {
      tmpH = n
    }

    if n < tmpL {
      tmpL = n
    }
  }
  return fmt.Sprintf("%d %d", tmpH, tmpL)
}
