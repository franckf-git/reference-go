package main

func insertionSort(datas []int) []int {
	datasLen := len(datas)
	for i := 1; i < datasLen; i++ {
		itemToSort := datas[i]
		indexOfPrevItem := i - 1
		for indexOfPrevItem >= 0 && datas[indexOfPrevItem] > itemToSort {
			datas[indexOfPrevItem+1] = datas[indexOfPrevItem]
			indexOfPrevItem--
		}
		datas[indexOfPrevItem+1] = itemToSort
	}
	return datas
}
