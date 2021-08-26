package main

import (
	"fmt"
	"reflect"
)

func main() {
	want := []int{1986, 1998, 2006, 2018, 2020}

	sortingB := bubbleSort([]int{2018, 1998, 1986, 2020, 2006})
	if reflect.DeepEqual(sortingB, want) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sortingB)
	}

	sortingI := insertionSort([]int{2018, 1998, 1986, 2020, 2006})
	if reflect.DeepEqual(sortingI, want) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sortingI)
	}

	sortingM := mergeSort([]int{2018, 1998, 1986, 2020, 2006})
	if reflect.DeepEqual(sortingM, want) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sortingM)
	}

	sortingQ := quickSort([]int{2018, 1998, 1986, 2020, 2006})
	if reflect.DeepEqual(sortingQ, want) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sortingQ)
	}
}
