// top

package kata

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

// ---

package kata

func RemoveChar(word string) (result string) {
  if len(word) > 2 {
    result = word[1 : len(word)-1]
  }
  return
}