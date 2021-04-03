package main

import "fmt"

func main() {

	// Creating slices from an array
	var arr1 = [4]int{1, 2, 3, 4}
	var arrSlice1 = arr1[1:3]
	fmt.Println(arrSlice1) //prints [2 3]

	// Creating slices from another slice
	var odds = [8]int{1, 3, 5, 7, 9, 11, 13, 15}
	slice1 := odds[2:]
	fmt.Println(slice1) // prints [5 7 9 11 13 15]
	slice2 := slice1[2:4]
	fmt.Println(slice2) // prints [9 11]

	// Creating slices using make() function
	s1 := make([]int, 5) // make an int slice of length 5
	fmt.Println(s1)      // [0 0 0 0 0]

	// The len() and cap() functions
	s2 := []int{2, 4, 6, 8, 10}
	fmt.Println(len(s2), cap(s2)) // 5 5
	arr2 := [6]int{1, 2, 3, 4, 5, 6}
	arrSlice2 := arr2[2:4]
	fmt.Println(len(arrSlice2), cap(arrSlice2)) // 2 4

	// Zero-valued slices
	s3 := []int{}
	fmt.Println(len(s3), cap(s3))

	// Modifying values in slices
	s4 := []int{1, 2, 3, 4}
	s4[1] = 0
	fmt.Println(s4) // [1 0 3 4]

	// Removing values from slices
	s5 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//v := []int{}
	_ = append(s5[:3], s5[4:]...) // remove 3rd elem
	fmt.Println(s5)
    // This statement drops the first and last elements of our slice:
    slice = slice[1:len(slice)-1]

	// Passing slice to a function
	s6 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	PrintSlice(s6) // [1 2 3 4 5 6 7 8]

	// Slice functions
	s7 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var s7b = make([]int, 5)
	// append function
	s7 = append(s7, 9, 10)
	fmt.Println(s7) // [1 2 3 4 5 6 7 8 9 10]
    // append two slices
    x := []int{1,2,3}
    y := []int{4,5,6}
    x = append(x, y...)
	// copy function
	ne := copy(s7b, s7)  // returns number of elements copied
	fmt.Println(s7b, ne) // [1 2 3 4 5] 5

	// Slice of slices
	mds := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 8}, // mandatory last comma
	}
	fmt.Println(mds) // [[1 2 3] [4 5 6]

	// Iterating over a slice elements
	s8 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range s8 {
		fmt.Println(i, v)
	}
	// output:
	// 0 1
	// 1 2
	// 2 3
	// 3 4
	// 4 5
	// 5 6
	// 6 7
	// 7 8
	// 8 9

}

func PrintSlice(a []int) {
	fmt.Println(a)
}

