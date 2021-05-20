package kata

func WordsToMarks(s string) int {
	result := 0
	arr := []rune(s)
	for _, char := range arr {
		result = int(char) - 96 + result
	}
	return result
}

// top

package kata

func WordsToMarks(s string) int {
  count := 0
  for _, i := range s {
     count += int(i) - 'a' + 1;
  }
  return count
}
