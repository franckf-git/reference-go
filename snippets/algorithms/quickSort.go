/*
def quick_sort(array_to_sort):
    # if there is no smaller or bigger values we have nothing to sort anymore
    # array is sorted !
    # this is the recursive stop condition
    if not array_to_sort:
        return []
    else:
        # if we have value to sort we choose a pivot
        # here we just getting the last element of the array
        pivot = array_to_sort[-1]

        # we go through the current array and build an array
        # of smaller values of the pivot
        array_of_smaller_values = [value for value in array_to_sort if value < pivot]

        # we go through the current array (minus the pivot) and build an array
        # of bigger values of the pivot
        array_of_bigger_values = [value for value in array_to_sort[:-1] if value >= pivot]

        # we return the current iteration with a new array
        # smaller values at the beggining, pivot in the middle, bigger values in the end
        return quick_sort(array_of_smaller_values) + [pivot] + quick_sort(array_of_bigger_values)

print(quick_sort([2018, 1998, 1986, 2020, 2006]))
# [1986, 1998, 2006, 2018, 2020]

*/
package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{2018, 1998, 1986, 2020, 2006})) // ? [1986, 1998, 2006, 2018, 2020]
}

func quickSort(datas []int) []int {

}
