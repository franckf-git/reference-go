package main

func Recursive(n int) uint64 {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Recursive(n-1) + Recursive(n-2)
	}
}
