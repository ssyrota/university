package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llinear_system
// #include "linear_system.h"
import "C"
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("hello from golang, platform: %s; arch: %s\n", runtime.GOOS, runtime.GOARCH)
	C.hello()
}
