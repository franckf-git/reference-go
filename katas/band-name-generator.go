package kata

import "strings"

func bandNameGenerator(word string) string {
	result := ""
	if word[0] == word[len(word)-1] {
		result = word + word[1:]
		return strings.Title(result)
	} else {
		result = "The " + word
		return strings.Title(result)
	}
}

// top

package kata

import "strings"

func bandNameGenerator(word string) string {
  first := word[:1]
  last := word[len(word)-1:]

  if first != last {
    return "The " + strings.Title(word)
  }

  return strings.Title(word) + word[1:]

}

// ---

package kata

import "strings"

func bandNameGenerator(word string) string {
  // Happy coding
  if word[0] == word[len(word)-1] {
    return strings.Title(word + word[1:])
  }
  return strings.Title("The " + word)
}

