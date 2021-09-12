package main

import "fmt"

type Queue struct {
	slice []int
}

func (q *Queue) Enqueue(add int) {
	q.slice = append(q.slice, add)
}

func (q *Queue) Dequeue() (first int) {
	if len(q.slice) == 0 {
		return -1
	}
	first = q.slice[0]
	q.slice = q.slice[1:len(q.slice)]
	return
}

func main() {
	var myMQ1 Queue
	// or like :
	var myMQ2 *Queue = new(Queue)
	myMQ1.Enqueue(1)
	myMQ1.Enqueue(1)
	myMQ2.Enqueue(10)
	myMQ2.Enqueue(5)
	myMQ2.Enqueue(6)
	myMQ2.Enqueue(2)
	myMQ2.Enqueue(1)
	myMQ1.Enqueue(1)
	myMQ2.Enqueue(2)
	myMQ2.Enqueue(1)
	myMQ2.Enqueue(8)

	fmt.Println(myMQ1.slice)
	fmt.Println(myMQ2)
	fmt.Println(myMQ2.slice)

	myMQ1.Dequeue()
	myMQ2.Dequeue()
	myMQ2.Dequeue()

	fmt.Println(myMQ1.slice)
	fmt.Println(myMQ2.slice)

	myMQ1.Dequeue()
	myMQ1.Dequeue()
	myMQ1.Dequeue()
	myMQ1.Dequeue()
	fmt.Println(myMQ1.Dequeue())
}
