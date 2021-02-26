package kata

func PositiveSum(numbers []int) int {
 var sum = 0
  for i := 0; i < len(numbers); i++ {
        if numbers[i] > 0 {
          sum = sum + numbers[i]
        }
    }
  return sum
}

// top

package kata

func PositiveSum(numbers []int) (sum int) {
  for _, num := range numbers {
    if num > 0 {
      sum = sum + num
    }
  }
  return
}

