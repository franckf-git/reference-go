package kata 

func Solve(arr []int) int {
  result := 0
	for _,v := range arr {
	for _,v2 := range arr {
        v = 0 - v
    if v != v2{
      result = 0 - v
    }
  }
  }
  return result
}
