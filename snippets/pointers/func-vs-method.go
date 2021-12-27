package main

import "log"

type Grid struct {
	num  [2][2]int
	mark [2][2]bool
}

// func with pointer call
//func punch(g *Grid, lin int, col int) {
//	g.mark[lin][col] = true
//}

// method
func (g *Grid) punch(lin, col int){
	g.mark[lin][col] = true
}

func main() {

	myG := Grid{num: [2][2]int{{5, 6}, {7, 8}}}
	log.Printf("myG: %#+v\n", myG)
	for il, l := range myG.num {
		for ic, c := range l {
			if c == 6 {

				//punch(&myG, il, ic)
                myG.punch(il,ic)
			}
		}
	}
	log.Printf("myG: %#+v\n", myG)
}
