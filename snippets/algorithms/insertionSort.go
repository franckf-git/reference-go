/*
unsorted_array = [2018, 1998, 1986, 2020, 2006]
def insertion_sort(array_to_sort):
    for index in range(1, len(array_to_sort)):
        current_item = array_to_sort[index]
        current_left_index = index - 1
        while current_left_index >= 0 and array_to_sort[current_left_index] > current_item:
            array_to_sort[current_left_index + 1] = array_to_sort[current_left_index]
            current_left_index -= 1
        array_to_sort[current_left_index + 1] = current_item
    return array_to_sort
print(insertion_sort(unsorted_array))
# [1986, 1998, 2006, 2018, 2020]
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	sorting := insertionSort([]int{2018, 1998, 1986, 2020, 2006})
	if reflect.DeepEqual(sorting, []int{1986, 1998, 2006, 2018, 2020}) {
		fmt.Println("ok")
	} else {
		fmt.Println("Nok")
	}
}

func insertionSort(datas []int) []int {
	return []int{}
}
