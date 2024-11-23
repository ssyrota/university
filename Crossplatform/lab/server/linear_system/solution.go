package linear_system

type Solution struct {
	A     [][]float64 `json:"a,omitempty"`
	B     []float64   `json:"b,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewErrorSolution(err string) Solution {
	return Solution{
		Error: err,
	}
}

func NewSolutionFromVector(vector []float64) Solution {
	length := len(vector)
	identityMatrix := make([][]float64, length)
	for i := 0; i < length; i++ {
		identityMatrix[i] = make([]float64, length)
		identityMatrix[i][i] = 1
	}
	return Solution{
		A: identityMatrix,
		B: vector,
	}
}
