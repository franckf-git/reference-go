// from https://www.calhoun.io/lets-learn-algorithms-graph-theory-representing-a-binary-tree-in-go-code/
package main

import "fmt"

// Node is used to represent a single vertex in our tree. 
// Each node has two children - a left child and a right 
// child. Either of these children can be set to nil which
// means that there is no child in that specific position.
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func main() {
	nodes := read()
	for _, node := range nodes {
		printNode(&node)
	}
}

func read() []Node {
	// First we declare and read in N, the total number of nodes in our graph
	var N int
	fmt.Scanf("%d", &N)
	
	// Next we create a slice of nodes of the size N, as we know we will be
	// reading in this many nodes from our input file.
	var nodes []Node = make([]Node, N)

	// Next we iterate over the N lines of input. On each line there are 3
	// integers. The first is the value of the node, the second is the index
	// of the left child, and the third is the index of the right child.
	//
	// Becuase we already instantiated our slice, we can freely set these nodes
	// to the left and right child of the current node even if we haven't set
	// the value of the node it is referencing. That will happen later in our
	// for loop.
	for i := 0; i < N; i++ {
		var val, indexLeft, indexRight int
		fmt.Scanf("%d %d %d", &val, &indexLeft, &indexRight)
		nodes[i].Value = val
		if indexLeft >= 0 {
			nodes[i].Left = &nodes[indexLeft]
		}
		if indexRight >= 0 {
			nodes[i].Right = &nodes[indexRight]
		}
	}
	
	// Finally we return our nodes :)
	return nodes
}

// Prints out a node's value plus the values of its children for debugging purposes.
func printNode(n *Node) {
	// Always print the value of the node
	fmt.Print("Value: ", n.Value)
	
	// If the node has a left child, print out it's value.
	if n.Left != nil {
		fmt.Print(" Left: ", n.Left.Value)
	}
	
	// If the node has a right child, print out it's value.
	if n.Right != nil {
		fmt.Print(" Right: ", n.Right.Value)
	}
	
	// Print out a newline character so our print statements aren't all on the same line.
	fmt.Println()
}

