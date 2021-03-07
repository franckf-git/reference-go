package kata

import "strings"

func RemoveChar(word string) string {
  wordarr := strings.Split(word, "")
  result := []string{}
  for key,value := range wordarr {
    if key != 0 && key != len(wordarr)-1 {
      newEntry := []string{value}
      result = append(result, newEntry...)
    }
  }
  return strings.Join(result, "")
}

// top

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
