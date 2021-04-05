package kata

func Grow(arr []int) int{
  var res int = 1
  for _, value := range arr {
    res = res * value
  }
  return res
}

// top

package kata

func Grow(arr []int) int{
  result := 1

  for _, n := range arr {
    result *= n
  }

  return result
}
