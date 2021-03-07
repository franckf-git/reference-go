package kata

import "strings"

func DNAStrand(dna string) string {
    re := strings.NewReplacer("A", "T", "C", "G", "T", "A", "G", "C")
         res := re.Replace(dna)
  return res
}

// top

package kata

import "strings"

var dnaReplacer *strings.Replacer = strings.NewReplacer(
  "A", "T",
  "T", "A",
  "C", "G",
  "G", "C",
)

func DNAStrand(dna string) string {
  return dnaReplacer.Replace(dna)
}


