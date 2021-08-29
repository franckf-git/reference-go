package main

func quickSort(datas []int) []int {
	if len(datas) == 0 {
		return []int{}
	}
	pivot := datas[len(datas)-1]
	smallerValues := []int{}
	largerValues := []int{}
	for index := 0; index < len(datas)-1; index++ {
		if datas[index] < pivot {
			smallerValues = append(smallerValues, datas[index])
		} else {
			largerValues = append(largerValues, datas[index])
		}
	}
	qsSmallerPivot := append(quickSort(smallerValues), pivot)
	return append(qsSmallerPivot, quickSort(largerValues)...)
}
