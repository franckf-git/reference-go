package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("filename missing")
		return
	}
	filename := os.Getenv("PWD") + "/" + os.Args[1]

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	output := tac(data)
	os.Stdout.Write([]byte(output))
}

func tac(in []byte) (out []byte) {
	var revert []string
	splitIn := strings.Split(string(in), "\n")
	for i := len(splitIn) - 1; i >= 0; i-- {
		revert = append(revert, splitIn[i])
	}
	out = []byte(strings.Join(revert, "\n"))
	return
}
