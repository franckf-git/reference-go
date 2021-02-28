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

// ---

func trimLeftChar(s string) string {
  for i := range s {
      if i > 0 {
          return s[i:]
      }
  }
  return s[:0]
}
