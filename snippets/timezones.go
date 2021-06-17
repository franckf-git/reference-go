package main

import (
	"fmt"
	"time"
)

func main() {
	// load time zone
	loc, e := time.LoadLocation("EST") // use other time zones such as MST, IST
    // loc, e := time.LoadLocation("America/Los_Angeles")
	CheckError(e)

	// get time in that zone
	fmt.Println(time.Now().In(loc)) // 2020-03-26 14:31:00.6625139 -0500 EST
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
