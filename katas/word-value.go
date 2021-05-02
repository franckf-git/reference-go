package kata

func NameValue(my_list []string) []int {
	var result = make([]int, len(my_list))
	for index, word := range my_list {

		wordR := []rune(word)
		addWord := 0
		for _, letter := range wordR {
			if letter == 32 {
				addWord = addWord
			} else {
				addWord = addWord + int(letter) - 96
			}
		}
		totalword := addWord * (index + 1)
		result[index] = totalword
	}
	return result
}

// top

package kata

func NameValue(my_list []string) []int {
  var result = make([]int, len(my_list))

  for idx, str := range my_list {
    for _, chr := range str {
      if chr >= 'a' && chr <= 'z' {
        result[idx] += int(chr-'a') + 1
      }
    }

    result[idx] = result[idx] * (idx + 1)
  }

  return result
}
