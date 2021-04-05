package kata

func ExpressionMatter(a int, b int, c int) int {
  res := 0
  fi := a * (b + c)
  se := a * b * c
  th := a + b * c
  qu := (a + b) * c
  fv := a + b + c
arr := [5]int{fi,se,th,qu,fv}
  for _, value := range arr {
    if value > res {
      res = value
    }
  }
  return res
}

// top

package kata

import "sort"

func ExpressionMatter(a int, b int, c int) int {
    arr := []int{a*(b+c), a*b*c, a+b+c, (a+b)*c}
    sort.Ints(arr)
    return arr[len(arr)-1]
}
