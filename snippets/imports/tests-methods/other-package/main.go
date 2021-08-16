package main

import (
	"ff/other-package/samelevel"
	"ff/other-package/sub-folder"
)

func main() {
	PublicFuncSLOP()
	//privateFuncSLOP()

	sub.PublicFuncSFOP()

	//sub.privateFuncSLOP() // undefined

	// sub folders and different package only work with go mod and if folder and package have the same name
}
