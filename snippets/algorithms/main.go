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

	wantBD := []int{24, 23, 22, 21, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	sortingBD := bubbleSortDesc([]int{21, 4, 2, 13, 10, 0, 19, 11, 7, 5, 23, 18, 9, 14, 6, 8, 1, 20, 17, 3, 16, 22, 24, 15, 12})
	if reflect.DeepEqual(sortingBD, wantBD) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sortingBD)
	}

	wantBS := []string{"alligator", "bald eagle", "bat", "camel", "cat", "cheetah", "colt", "cow", "dog", "dung beetle", "frog", "hamster", "horse", "mink", "moose", "porcupine", "rat", "rooster", "steer"}
	sortingBS := bubbleSortStrings([]string{"dog", "cat", "alligator", "cheetah", "rat", "moose", "cow", "mink", "porcupine", "dung beetle", "camel", "steer", "bat", "hamster", "horse", "colt", "bald eagle", "frog", "rooster"})
	if reflect.DeepEqual(sortingBS, wantBS) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(sortingBS)
	}

	wantBinS := 3
	sortedList := []int{1, 3, 4, 6, 7, 9, 10, 11, 13}
	gotBinS := binarySearch(sortedList, 6)
	if wantBinS == gotBinS {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(gotBinS)
	}
	gotBinS2 := binarySearch2(sortedList, 6)
	if wantBinS == gotBinS2 {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
		fmt.Println(gotBinS2)
	}
}
