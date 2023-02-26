package main

import (
	"fmt"
	"runtime"
)

var version = "dev"

func main() {
	fmt.Println("Hello World! (", version, runtime.GOOS, runtime.GOARCH, ")")
}
