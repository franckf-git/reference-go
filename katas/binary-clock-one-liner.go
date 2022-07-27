package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		tNow := time.Now()
		tHour, tMins, tSecs := tNow.Clock()
		bHour := fmt.Sprintf("%b", tHour)
		bMins := fmt.Sprintf("%b", tMins)
		bSecs := fmt.Sprintf("%b", tSecs)
		time.Sleep(time.Second)
		fmt.Printf("h:%v m:%v s:%v           \r", bHour, bMins, bSecs)
	}
}
