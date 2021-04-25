package kata

func InAscOrder(numbers []int) bool {
	if len(numbers) == 1 {
		return true
	}
	for ind := 0; ind < len(numbers)-1; ind++ {
// not great - doesn't compare the last value
		if numbers[ind] > numbers[ind+1] {
			return false
		}
	}
	return true

}

// top

package kata

import (
  "sort"
)

func InAscOrder(numbers []int) bool {
  return sort.IntsAreSorted(numbers)
}
