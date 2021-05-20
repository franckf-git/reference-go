// Tail recursion happens when a function calling itself calculates the value and sends it down the hierarchy. Here is an example of a tail-recursion.

// The tail-calls are beneficial in an aspect that it is optimized by compilers in many cases. So, it is almost always faster than a normal recursive call.

package main

import (
    "fmt"
)

func f(v int) {
    if(v == 0) {
        fmt.Println("Zero")
        return
    } else {
        fmt.Println(v)
        f(v-1)               // tail-recursive call
    }
}

func main() {
    f(5)
    // output:
    // 5
    // 4
    // 3
    // 2
    // 1
    // Zero
}
