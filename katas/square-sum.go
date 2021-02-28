package kata

func SquareSum(numbers []int) int {
	total := 0
	for _, v := range numbers {
		p := v * v
		total = total + p
	}
	return total
}

// top

package kata

func SquareSum(nums []int) (res int) {
    for _, val := range nums {
      res += val * val
    }
    return res
}