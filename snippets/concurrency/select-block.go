package main

import "time"
import "fmt"

func main() {
    c := make(chan string)
    q := make(chan bool)
    go myDelayedQuit(q)
    go mySleep(c, 1)

    fmt.Println("begin non blocking wait...")
    for {
        select {
        case msg := <-c:
            fmt.Println("received:", msg)
      case <-q:
            fmt.Println("done")
            return
        }
    }
}

func mySleep(a chan string, n int) {
    time.Sleep(time.Second * time.Duration(n))
    a <- "woke up"
}

func myDelayedQuit(b chan bool) {
    time.Sleep(time.Second * 2)
    b <- true
}
// for loops forever until the return statement
// select will wait and whenever a case can be filled it will unblock
// after 1 second the sleep function is done and sends the "woke up" message
// after 2 seconds the true boolean is sent and main finishes
