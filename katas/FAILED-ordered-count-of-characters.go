// top

package orderedcount

import "strings"

func OrderedCount(text string) []Tuple {
  counts := make([]Tuple, 0, len(text))
  counted := ""

  for _, r := range(text) {
    if strings.Count(counted, string(r)) == 0 {
      counts = append(counts, Tuple{Char: r, Count: strings.Count(text, string(r))})
      counted += string(r)
    }
  }
  return counts
}
