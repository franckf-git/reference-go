// Regular recursion is when calling itself does not finish the calculation immediately instead it passes it down the hierarchy. Below is an example showing that.

package main

import (
    "fmt"
)

func fib(n int) int {
    if (n == 0) {
        return 0
    } else if(n == 1) {
        return 1
    } else {
        return fib(n-1) + fib(n-2)
    }
}

func main() {
    fmt.Println(fib(10))  // 55
}
