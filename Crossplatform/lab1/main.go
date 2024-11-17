package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

// #cgo CFLAGS: -I.
// #cgo LDFLAGS: -L. -llinear_system
// #include "linear_system.h"
// #include <stdlib.h>
// #include <stdio.h>
import "C"

func main() {
	fmt.Printf("Hello from golang, platform: %s; arch: %s\n", runtime.GOOS, runtime.GOARCH)
	C.hello()

	ls := NewLinearSystem([][]int{{4, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3})
	defer ls.Free()
	C.print_vector(ls.b)
	C.print_matrix(ls.A)

	fmt.Println(ls.SolveGauss())
	// fmt.Println(ls.SolveMatrix())
}

func NewLinearSystem(a [][]int, b []int) *LinearSystem {
	matrix, freeMatrix := MakeMatrix(a)
	vector, freeVector := MakeVector(b)
	return &LinearSystem{
		A: matrix,
		b: vector,
		free: func() {
			freeMatrix()
			freeVector()
		},
	}
}

type LinearSystem struct {
	A    *C.Matrix
	b    *C.Vector
	free func()
}

func (ls *LinearSystem) Free() {
	ls.free()
}

func (ls *LinearSystem) SolveGauss() [][]int {
	matrix := C.make_AugmentedMatrix(ls.A, ls.b)
	linearSystem := C.make_LinearEquationSystem(matrix)
	augmentedMatrix := C.LinearEquationSystem_solve_gauss(linearSystem)
	return ParseMatrix(augmentedMatrix.matrix)
}

func (ls *LinearSystem) SolveMatrix() [][]int {
	matrix := C.make_AugmentedMatrix(ls.A, ls.b)
	linearSystem := C.make_LinearEquationSystem(matrix)
	augmentedMatrix := C.LinearEquationSystem_solve_matrix(linearSystem)
	return ParseMatrix(augmentedMatrix.matrix)
}

func MakeVector(a []int) (*C.Vector, func()) {
	length := len(a)
	cArr := make([]C.int, length)
	for i, v := range a {
		cArr[i] = C.int(v)
	}
	cVec := C.make_Vector(&cArr[0], C.int(length))
	return cVec, func() {
		C.free(unsafe.Pointer(cVec))
	}
}

func ParseVector(cVec *C.Vector) []int {
	length := int((*cVec).length)
	a := make([]int, length)
	for i := 0; i < length; i++ {
		a[i] = int(*(*C.int)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cVec.a)) + uintptr(C.sizeof_int*i)),
		))
	}
	return a
}

func MakeMatrix(a [][]int) (*C.Matrix, func()) {
	rows := len(a)
	cols := len(a[0])
	cArr := make([]*C.Vector, rows)
	for i, row := range a {
		vector, _ := MakeVector(row)
		cArr[i] = vector
	}
	cMat := C.make_Matrix((&cArr[0]), C.int(rows), C.int(cols))
	return cMat, func() {
		for _, vector := range cArr {
			C.free(unsafe.Pointer(vector))
		}
		C.free(unsafe.Pointer(cMat))
	}
}

/*
The problem is that pointer in C can be treated as pointer to struct and as an array.
So the matrix have **Vector attribute, which should be treated as an array of *Vector.
*/
func ParseMatrix(cMat *C.Matrix) [][]int {
	rows := int(cMat.rows)
	a := make([][]int, rows)

	vectorArrayPtr := uintptr(unsafe.Pointer((*cMat).a))
	pointerSize := uintptr(unsafe.Sizeof((*cMat).a))
	for i := 0; i < rows; i++ {
		vector := (**C.Vector)(unsafe.Pointer(
			uintptr(vectorArrayPtr + pointerSize*uintptr(i)),
		))
		a[i] = ParseVector(*vector)
	}
	return a
}
