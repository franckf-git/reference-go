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

func bubbleSortDesc(datas []int) []int {
	datasLen := len(datas)
	for i := 0; i < datasLen; i++ {
		sweep(datas)
	}
	return datas
}

func sweep(numbers []int) {
	numbersLen := len(numbers)
	number1 := 0
	number2 := 1

	for number2 < numbersLen {
		compare1 := numbers[number1]
		compare2 := numbers[number2]
		if compare1 < compare2 {
			numbers[number1] = compare2
			numbers[number2] = compare1
		}
		number1++
		number2++
	}
}

func bubbleSortStrings(datas []string) []string {
	datasLen := len(datas)
	for i := 0; i < datasLen; i++ {
		swapString(datas)
	}
	swapString(datas)
	return datas
}

func swapString(words []string) {
	wordsLen := len(words)
	word1 := 0
	word2 := 1

	for word2 < wordsLen {
		compare1 := words[word1]
		compare2 := words[word2]
		if compare1 > compare2 {
			words[word1] = compare2
			words[word2] = compare1
		}
		word1++
		word2++
	}
}

/*
 * type Interface interface {
 * Len() int
 * Less(i, j int) bool
 * Swap(i, j int)
 * }
 */
func bubbleSortInterface(x sort.Interface) {
	n := x.Len()
	for {
		swapped := false
		for i := 1; i < n; i++ {
			if x.Less(i, i-1) {
				x.Swap(i, i-1)
				swapped = true
			}
		}
		if !swapped {
			return
		}
	}
}
