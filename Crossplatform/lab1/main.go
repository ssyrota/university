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

	ls1 := NewLinearSystem([][]float64{
		{1, 1, 1},
		{4, 2, 1},
		{9, 3, 1},
	}, []float64{0, 1, 3})
	defer ls1.Free()
	C.print_vector(ls1.b)
	C.print_matrix(ls1.A)
	fmt.Println("SolveMatrix: ", ls1.SolveMatrix())

}

func NewLinearSystem(a [][]float64, b []float64) *LinearSystem {
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

func (ls *LinearSystem) SolveGauss() [][]float64 {
	matrix := C.make_AugmentedMatrix(ls.A, ls.b)
	linearSystem := C.make_LinearEquationSystem(matrix)
	augmentedMatrix := C.LinearEquationSystem_solve_gauss(linearSystem)
	return ParseMatrix(augmentedMatrix.matrix)
}

func (ls *LinearSystem) SolveMatrix() []float64 {
	matrix := C.make_AugmentedMatrix(ls.A, ls.b)
	linearSystem := C.make_LinearEquationSystem(matrix)
	vector := C.LinearEquationSystem_solve_matrix(linearSystem)
	if vector == nil {
		fmt.Println("Linear system not solvable by matrix method")
		return nil
	}
	return ParseVector(vector)
}

func MakeVector(a []float64) (*C.Vector, func()) {
	length := len(a)
	cArr := make([]C.float, length)
	for i, v := range a {
		cArr[i] = C.float(v)
	}
	cVec := C.make_Vector(&cArr[0], C.int(length))
	return cVec, func() {
		C.free(unsafe.Pointer(cVec))
	}
}

func ParseVector(cVec *C.Vector) []float64 {
	length := int((*cVec).length)
	a := make([]float64, length)
	for i := 0; i < length; i++ {
		a[i] = float64(*(*C.float)(unsafe.Pointer(
			uintptr(unsafe.Pointer(cVec.a)) + uintptr(C.sizeof_float*i),
		)))
	}
	return a
}

func MakeMatrix(a [][]float64) (*C.Matrix, func()) {
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
func ParseMatrix(cMat *C.Matrix) [][]float64 {
	rows := int(cMat.rows)
	a := make([][]float64, rows)

	vectorArrayPtr := uintptr(unsafe.Pointer((*cMat).row_vectors))
	pointerSize := uintptr(unsafe.Sizeof((*cMat).row_vectors))
	for i := 0; i < rows; i++ {
		vector := (**C.Vector)(unsafe.Pointer(
			uintptr(vectorArrayPtr + pointerSize*uintptr(i)),
		))
		a[i] = ParseVector(*vector)
	}
	return a
}
