package main

import "fmt"

func main() {
	i := 1
	nopoit(i)
	fmt.Println(i) // 1
    j := 1
	poit(&j)
	fmt.Println(j) // 2
}

func poit(num *int) {
	*num++
}

func nopoit(num int) {
	num++
}
