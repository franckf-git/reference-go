package kata

func RepeatStr(repetitions int, value string) string {
	result := ""
	for i := 0; i < repetitions; i++ {
		result = result + string(value)
	}
	return result
}

// top

package kata

import "strings"

func RepeatStr(repititions int, value string) string {
  return strings.Repeat(value, repititions)
}

// ---

package kata

func RepeatStr(repititions int, value string) (result string) {
  for i:=0; i < repititions; i++ {
      result += value
  }
  return
}