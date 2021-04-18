

// top

package kata
import (
  "strings"
  "strconv"
)

func NbDig(n int, d int) (out int) {
  for i := 0; i <= n; i++ {
    out += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
  }

  return
}

// ---

package kata

func NbDig(n int, d int) int {
    if n == 0 {
        return 0
    }

    count := 0
    for i := 0; i <= n; i++ {
        count += digits(i*i, d)
    }
    return count
}

func digits(val, d int) int {
    if val == d {
        return 1
    }

    count := 0
    for val != 0 {
        digit := val % 10
        if digit == d {
          count++
        }
        val /= 10
    }
    return count
}
