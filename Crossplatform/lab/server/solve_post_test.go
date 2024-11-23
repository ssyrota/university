package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSolvePost(t *testing.T) {
	type testcase struct {
		input    any
		expected any
		name     string
	}
	testcases := []testcase{
		{
			input: map[string]any{
				"a": [][]float64{
					{4, 7},
					{2, 6},
				},
				"b": []float64{15, 10},
			},
			expected: map[string]any{
				"a": [][]float64{
					{1, 0},
					{0, 1},
				},
				"b": []float64{2, 1},
			},
			name: "2x2 matrix solvable",
		},
		{
			input: map[string]any{
				"a": [][]float64{
					{1, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 0},
					{0, 0, 0, 0, 1},
				},
				"b": []float64{1, 2, 3, 4, 5},
			},
			expected: map[string]any{
				"a": [][]float64{
					{1, 0, 0, 0, 0},
					{0, 1, 0, 0, 0},
					{0, 0, 1, 0, 0},
					{0, 0, 0, 1, 0},
					{0, 0, 0, 0, 1},
				},
				"b": []float64{1, 2, 3, 4, 5},
			},
			name: "5x5 matrix solvable",
		},
		{
			input: map[string]any{
				"a": [][]float64{
					{1, 2},
					{2, 4},
				},
				"b": []float64{3, 6},
			},
			expected: map[string]any{
				"error": "matrix is singular",
			},
			name: "2x2 singular matrix",
		},
	}

	for _, tc := range testcases {
		jsonBody, _ := json.Marshal(tc.input)
		req := httptest.NewRequest(http.MethodPost, "/solve", io.NopCloser(bytes.NewReader(jsonBody)))
		w := httptest.NewRecorder()
		SolvePost(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("[%s] expected error to be nil got %v", tc.name, err)
		}
		jsonExpected, _ := json.Marshal(tc.expected)

		received := strings.TrimSpace(string(data))
		expected := strings.TrimSpace(string(jsonExpected))
		if received != expected {
			t.Errorf("[%s] expected %v got %v", tc.name, expected, received)
		}
	}
}
