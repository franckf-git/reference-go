// if the sender create more that the receiver can get

package main

import (
    "fmt"
    "log"
    "sync"
    "time"
)

func slowReceiver(c <-chan int, wg *sync.WaitGroup) {
    for {
        n := <-c
        fmt.Println("received", n)
        time.Sleep(time.Second)
        // c <- 42 // THIS WOULD CAUSE AN ERROR "(send to receive-only type <-chan int)"
        wg.Done()
    }
}

func fastSender(c chan<- int, wg *sync.WaitGroup) {
    for i := 0; i < 5; i++ {
        c <- i
        wg.Add(1)
        fmt.Println(i)
        // fmt.Println(<-c) // THIS WOULD CAUSE AN ERROR "(receive from send-only type chan<- int)"
    }
}

func main() {
    start := time.Now()
    var wg sync.WaitGroup
    c := make(chan int, 3)
    go slowReceiver(c, &wg)
    fastSender(c, &wg)
    wg.Wait()
    log.Println(time.Since(start))
}
