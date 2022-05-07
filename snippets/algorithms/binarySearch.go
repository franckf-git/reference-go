package main

// binarySearch calhoun implement no recursif and no new slice
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

// binarySearch2 my implement with recurvive
func binarySearch2(numbers []int, search int) int {
	var midIndex int = len(numbers) / 2
	var midVal int = numbers[midIndex]
	if midVal == search {
		return midIndex
	}
	var left []int = numbers[:midIndex]
	var righ []int = numbers[midIndex+1:]
	if midVal > search {
		return binarySearch2(left, search)
	} else {
		return binarySearch2(righ, search) + midIndex + 1
	}
}

// binarySearch(numbers, 0, len(numbers)-1, numberToFind)
func binarySearchOther(numbers []int, leftBound, rightBound, numberToFind int) int {
	if rightBound >= leftBound {
		midPoint := leftBound + (rightBound-leftBound)/2

		if numbers[midPoint] == numberToFind {
			return midPoint
		}

		if numbers[midPoint] > numberToFind {
			return binarySearch(numbers, leftBound, midPoint-1, numberToFind)
		}

		return binarySearch(numbers, midPoint+1, rightBound, numberToFind)
	}

	return -1
}
