package main

func bubbleSort(datas []int) []int {
	datasLen := len(datas)
	for i := 0; i < datasLen; i++ {
		if !swap(datas, i) {
			return datas
		}
	}
	return datas
}

func swap(numbers []int, prevPass int) bool {
	numbersLen := len(numbers)
	number1 := 0
	number2 := 1
	swapDone := false
	for number2 < numbersLen-prevPass {
		compare1 := numbers[number1]
		compare2 := numbers[number2]
		if compare1 > compare2 {
			numbers[number1] = compare2
			numbers[number2] = compare1
			swapDone = true
		}
		number1++
		number2++
	}
	return swapDone
}

// TODO do calhoun execises (desc, strings, names)

func bubbleSortDesc(datas []int) []int {
	return []int{}
}

func bubbleSortStrings(datas []string) []string {
	return []string{}
}
