package main

import (
	"fmt"
	"math/rand"

	"github.com/sergeycrisp/university/RTPP/plot"
	qh "github.com/sergeycrisp/university/RTPP/quick_hull"
)

func main() {
	all := GeneratePoints(100)
	corner, allPoints := qh.QuickHull(all)
	plot.MakePlot(allPoints, corner, "2.png")
}

//Generate points
func GeneratePoints(n int) []qh.Point {
	res := make([]qh.Point, n)
	for i := 0; i < n; i++ {
		res[i] = qh.Point{rand.Float64() * float64(n), rand.Float64() * float64(n), fmt.Sprint(i)}
	}
	return res
}
