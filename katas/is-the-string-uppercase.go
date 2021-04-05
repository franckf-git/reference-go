package kata
import "strings"
type MyString string

func (s MyString) IsUpperCase() bool {
  upper := strings.ToUpper(string(s))
  return string(s) == upper
}

// top

package kata

import (
  "strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
  return s == MyString(strings.ToUpper(string(s)))
}

// ---

package kata

type MyString string

func (s MyString) IsUpperCase() bool {
  for _, letter := range s {
    if int(letter) >= 97 && int(letter) <= 122 {
      return false
    }
  }
  return true
}
