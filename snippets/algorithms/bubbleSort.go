package main

import (
	"fmt"
	"reflect"
)

func main() {
	sorting := bubbleSort([]int{2018, 1998, 1986, 2020, 2006})
	if reflect.DeepEqual(sorting, []int{1986, 1998, 2006, 2018, 2020}) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sorting)
	}
}

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

// TODO try and benchmark longer slices with and without optimize

// TODO do calhoun execises (desc, strings, names)
