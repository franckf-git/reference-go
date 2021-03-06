package kata

import (
	"sort"
	"strings"
)

func TwoToOne(s1 string, s2 string) string {
	// combine and sort strings
	twoS := s1 + s2
	twoSslice := strings.Split(twoS, "")
	sort.Strings(twoSslice)

	// create a map to record if a value is already here
	occured := map[string]bool{}
	uniq := []string{}

	for key := range twoSslice {
		if occured[twoSslice[key]] != true {
			occured[twoSslice[key]] = true
			uniq = append(uniq, twoSslice[key])
		}
	}
	return strings.Join(uniq, "")
}

// top
package kata

import (
  "sort"
  "strings"
)

func TwoToOne(s1 string, s2 string) string {
  chars := strings.Split(s1 + s2, "")
  sort.Strings(chars)
  result := ""
  for _, s := range chars {
    chr := string(s)
    if !strings.Contains(result, chr) {
      result = result + chr
    }
  }
  return result
}