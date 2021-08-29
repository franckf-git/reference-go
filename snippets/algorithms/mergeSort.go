package main

import "fmt"

func mergeSort(datas []int) []int {
	if len(datas) < 2 {
		return datas
	}
	middle := len(datas) / 2
	fmt.Println(datas, middle)
	leftSide := mergeSort(datas[:middle])
	rightSide := mergeSort(datas[middle:])
	return merge(leftSide, rightSide)
}

func merge(leftSide []int, rightSide []int) []int {
	result := []int{}
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(leftSide) && rightIndex < len(rightSide) {
		if leftSide[leftIndex] < rightSide[rightIndex] {
			result = append(result, leftSide[leftIndex])
			leftIndex++
		} else {
			result = append(result, rightSide[rightIndex])
			rightIndex++
		}
	}
	for leftIndex < len(leftSide) {
		result = append(result, leftSide[leftIndex])
		leftIndex++
	}
	for rightIndex < len(rightSide) {
		result = append(result, rightSide[rightIndex])
		rightIndex++
	}
	return result
}
