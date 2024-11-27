package linear_system

// #cgo CFLAGS: -I../../c_library
// #cgo darwin,arm64 LDFLAGS: -L../../c_library/bin -ldarwin_arm64_linear_system
// #cgo linux,amd64 LDFLAGS: -L../../c_library/bin -llinux_amd64_linear_system
// #cgo linux,arm64 LDFLAGS: -L../../c_library/bin -llinux_arm64_linear_system
// #include "linear_system.h"
// #include <stdlib.h>
// #include <stdio.h>
import "C"

import (
	"runtime"
	"unsafe"
)

var pinner = new(runtime.Pinner)

func NewLinearSystem(a [][]float64, b []float64) *LinearSystem {
	return &LinearSystem{
		A: a,
		B: b,
	}
}

type LinearSystem struct {
	A [][]float64 `json:"a"`
	B []float64   `json:"b"`
}

func (ls *LinearSystem) SolveMatrix() Solution {
	matrixA, freeMatrixA := MakeMatrix(ls.A)
	vectorB, freeVectorB := MakeVector(ls.B)
	defer freeMatrixA()
	defer freeVectorB()

	matrix := C.make_AugmentedMatrix(matrixA, vectorB)
	linearSystem := C.make_LinearEquationSystem(matrix)
	vector := C.LinearEquationSystem_solve_matrix(linearSystem)
	if vector == nil {
		return NewErrorSolution("matrix is singular")
	}
	return NewSolutionFromVector(ParseVector(vector))
}

func (ls *LinearSystem) Determinant() float64 {
	matrixA, freeMatrixA := MakeMatrix(ls.A)
	defer freeMatrixA()
	cfloat := C.Matrix_determinant(matrixA)
	return float64(cfloat)
}

func MakeVector(a []float64) (*C.Vector, func()) {
	length := len(a)
	cArr := make([]C.float, length)
	pinner.Pin(&cArr)
	for i, v := range a {
		cArr[i] = C.float(v)
	}
	cVec := C.make_Vector(&cArr[0], C.int(length))
	return cVec, func() {
		C.free(unsafe.Pointer(cVec))
		pinner.Unpin()
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
