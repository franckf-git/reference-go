package kata
import "strings"
import "strconv"

func Points(games []string) int {
  var point int

  for _,value := range games {
      sub := strings.Split(value, ":")
    x,_ := strconv.Atoi(sub[0])
    y,_ := strconv.Atoi(sub[1])
    if x > y {
      point += 3
    }
    if x == y {
      point += 1
    }
  }
  return point;
}

// top

package kata

import ( "strings" )

func Points(games []string) int {
  result := 0
  for _, game := range games {
    str := strings.Split(game, ":")
    x, y := str[0], str[1]
    switch {
      case x > y:
        result += 3
      case x == y:
        result += 1
    }
  }
  return result
}
