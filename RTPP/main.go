package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	qh "github.com/sergeycrisp/university/RTPP/quick_hull"
)

func main() {
	for _, proc := range procs {
		for _, point := range pointsSorted {
			fmt.Printf("Input: %v Procs: %v \n", point, proc)
			runtime.GOMAXPROCS(proc)
			start := time.Now()
			qh.QuickHull(points[point])
			t := time.Now()
			elapsed := t.Sub(start)
			fmt.Println("Function_time: ", elapsed.Milliseconds(), " ms")
			fmt.Println()
		}
	}

	// plot.MakePlot(allPoints, corner, "2.png")
}

//Generate points
func GeneratePoints(n int) []qh.Point {
	res := make([]qh.Point, n)
	for i := 0; i < n; i++ {
		res[i] = qh.Point{X: rand.Float64() * float64(n), Y: rand.Float64() * float64(n), Name: fmt.Sprint(i)}
	}
	return res
}

var points = map[int][]qh.Point{
	1000:    GeneratePoints(1000),
	2000:    GeneratePoints(2000),
	10000:   GeneratePoints(10000),
	100000:  GeneratePoints(100000),
	1000000: GeneratePoints(1000000),
}
var pointsSorted = []int{1000, 2000, 10000, 100000, 1000000}

var procs = []int{1, 2, 4, 8}
