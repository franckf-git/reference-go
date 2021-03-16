package kata

func Maps(x []int) []int {
result :=  []int{}
  for _,value := range x {
  result = append(result, value*2)
}
  return result
}

// top
package kata

func Maps(x []int) []int {
  result := make([]int, len(x))
  for i, v := range x {
    result[i] = v + v
  }
  return result
}
