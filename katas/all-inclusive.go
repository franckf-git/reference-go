package kata
import (
    "sort"
    "strings"
)
func ContainAllRots(strng string, arr []string) bool {
   if strng == "" {
     return true
   }
  sortIn := strings.Split(strng, "")
  sort.Strings(sortIn)
  countOK := 0
  for _,v := range arr {
    sortV := strings.Split(v, "")
    sort.Strings(sortV)
    if strings.Join(sortIn, "") == strings.Join(sortV, "") {
      countOK = countOK + 1
    }
  }
  if countOK == len(strng) {
    return true
  }
  return false
}
