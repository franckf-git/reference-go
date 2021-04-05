package kata

import "strings"

func DNAtoRNA(dna string) string {
  re := strings.NewReplacer("T", "U")
  res := re.Replace(dna)
  return res
}

// top

package kata

import (
  "strings"
)

func DNAtoRNA(dna string) string {
  return strings.Replace(dna, "T", "U", -1)
}
