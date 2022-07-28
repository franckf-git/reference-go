package main

import "fmt"

/*
procédure Hanoï(n, D, A, I)
    si n ≠ 0
        Hanoï(n-1, D, I, A)
        Déplacer le disque de D vers A
        Hanoï(n-1, I, A, D)
    fin-si
fin-procédure
*/

func main() {
	towerA := []int{3, 2, 1}
	towerB := []int{}
	towerC := []int{}
	var moving int
	move(3, towerA, towerC, towerB, &moving)
	fmt.Println(moving)
}

func move(disks int, source, target, aux []int, moving *int) {
	if disks > 0 {
		move(disks-1, source, aux, target, moving)
		if len(source) > 0 {
			target = append(target, source[len(source)-1])
			source = source[:len(source)-1]
		}
		*moving++
		move(disks-1, aux, target, source, moving)
	}
}
