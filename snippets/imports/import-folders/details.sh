cd ~
mkdir -p import-folders/misc
cd import-folders
touch main.go
cd misc
touch whatever.go
cd ..
go mod init gitlab.com/ff/import-folders

In main.go:
package main

import "gitlab.com/ff/import-folders/misc"

func main() {
    misc.Hello()
}

In misc/whatever.go:
package misc

import "fmt"

func Hello() {
    fmt.Println("hello")
}
