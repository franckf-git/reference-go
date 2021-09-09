package main

// binarySearch clahoun implement no recursif and no new slice
func binarySearch(numbers []int, search int) int {
	lo := 0
	hi := len(numbers) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		midValue := numbers[mid]

		if midValue == search {
			return mid
		} else if midValue > search {
			// We want to use the left half of our list
			hi = mid - 1
		} else {
			// We want to use the right half of our list
			lo = mid + 1
		}
	}

	// If we get here we tried to look at an invalid sub-list
	// which means the number isn't in our list.
	return -1
}
