package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/alexeyco/simpletable"
	qh "github.com/sergeycrisp/university/RTPP/quick_hull"
)

func main() {
	// Benchmark(qh.QuickHullSequential, "Sequential")
	// Benchmark(qh.QuickHullMpi, "MPI")
	Benchmark(qh.QuickHullOpenMp, "OpenMP")
}

func Benchmark(f func(points []qh.Point) []qh.Point, name string) {
	fmt.Println(fmt.Sprintf("\033[36m%"+"17"+"s \033[90m:\033[0m ", name))

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Input"},
			{Align: simpletable.AlignCenter, Text: "Procs"},
			{Align: simpletable.AlignCenter, Text: "Time(ms)"},
		},
	}

	for _, proc := range procs {
		for _, point := range pointsSorted {
			runtime.GOMAXPROCS(proc)
			start := time.Now()
			f(points[point])
			t := time.Now()
			elapsed := t.Sub(start)

			r := []*simpletable.Cell{
				{Text: fmt.Sprint(point)},
				{Text: fmt.Sprint(proc)},
				{Text: fmt.Sprint(elapsed.Milliseconds())},
			}
			table.Body.Cells = append(table.Body.Cells, r)
		}
	}
	table.SetStyle(simpletable.StyleDefault)
	fmt.Println(table.String())
	fmt.Println()
}

//Generate points
func GeneratePoints(n int) []qh.Point {
	res := make([]qh.Point, n)
	for i := 0; i < n; i++ {
		res[i] = qh.Point{X: rand.Float64() * float64(n), Y: rand.Float64() * float64(n), Name: i}
	}
	return res
}

var points = map[int][]qh.Point{
	1000:    qh.EnumPossiblePoints(GeneratePoints(1000)),
	10000:   qh.EnumPossiblePoints(GeneratePoints(10000)),
	100000:  qh.EnumPossiblePoints(GeneratePoints(100000)),
	5000000: qh.EnumPossiblePoints(GeneratePoints(5000000)),
}
var pointsSorted = []int{1000, 10000, 100000, 5000000}

var procs = []int{1, 2, 4, 8}
