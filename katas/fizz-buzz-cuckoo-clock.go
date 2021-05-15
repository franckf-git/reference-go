package main

import "strconv"

func FizzBuzzCuckooClock(time string) string {
	hour, _ := strconv.Atoi(time[:2])
	min, _ := strconv.Atoi(time[3:])
	var result string

	if min == 00 {
		if hour > 12 {
			hour = hour - 12
		}
		if hour == 00 {
			hour = 12
		}
		for i := 1; i <= hour; i++ {
			result += " Cuckoo"
		}
		result = result[1:]
		return result
	} else if min == 30 {
		result += "Cuckoo"
		return result
	} else if min%5 == 0 && min%3 == 0 {
		result = "Fizz Buzz"
	} else if min%3 == 0 {
		result = "Fizz"
	} else if min%5 == 0 {
		result = "Buzz"
	} else {
		result = "tick"
	}

	return result
}

// top

package kata
import (
  "strings"
  "strconv"
)
func FizzBuzzCuckooClock(time string) string {
  h, _ := strconv.Atoi(time[0:2])
  m, _ := strconv.Atoi(time[3:5])
  switch {
    case m==0:
      return strings.Repeat("Cuckoo ", (h+11)%12) + "Cuckoo"
    case m==30:
      return "Cuckoo"
    case m%15 == 0:
      return "Fizz Buzz"
    case m%5 == 0:
      return "Buzz"
    case m%3 == 0:
      return "Fizz"
    default:
      return "tick"
  }
}

// ---

package kata

import (
  "strconv"
  "strings"
)

func FizzBuzzCuckooClock(time string) string {
  timeSplit := strings.Split(time, ":")
  hour, _ := strconv.Atoi(timeSplit[0])
  min, _ := strconv.Atoi(timeSplit[1])
  if min == 0 {
    if hour > 12 {
      hour -= 12
    }
    if hour == 0 {
      hour = 12
    }
    return strings.Repeat("Cuckoo ", hour-1) + "Cuckoo"
  } else if min == 30 {
    return "Cuckoo"
  } else if (min % 15) == 0 {
    return "Fizz Buzz"
  } else if (min % 3) == 0 {
    return "Fizz"
  } else if (min % 5) == 0 {
    return "Buzz"
  }
  return "tick"
}
