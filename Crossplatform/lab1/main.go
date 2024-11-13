package main

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llinear_system
// #include "linear_system.h"
import "C"

func main() {
	C.hello()
}
