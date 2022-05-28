package main

import (
	"fmt"

	"github.com/Arafatk/glot"
	"github.com/elliotchance/pie/v2"
	"golang.org/x/exp/constraints"
)

func main() {
	all := []Point{{1, 2, "1"}, {2, 1, "2"}, {4, 3, "3"}, {5, 3, "4"}, {6, 1, "5"}, {7, 4, "6"}, {8, 2, "7"}, {9, 3, "8"}, {10, 3, "9"}}

	// MakePlot(all, all, "2.png")
	fmt.Print(all[2].SideByLine(all[0], all[5]))
}

// Plot point
type Point struct {
	X    float64
	Y    float64
	Name string
}

// Draw plot to the file
func MakePlot(shellAllPoints, shellCornedPoints []Point, output string) error {
	// Set up plot
	dimensions := 2
	persist := false
	debug := false
	plot, err := glot.NewPlot(dimensions, persist, debug)
	if err != nil {
		return err
	}

	// Draw all points
	allPoints := [][]float64{{}, {}}
	for _, p := range shellAllPoints {
		allPoints[0] = append(allPoints[0], p.X)
		allPoints[1] = append(allPoints[1], p.Y)
	}
	plot.AddPointGroup("All points", "points", allPoints)

	// Draw corner shell points line
	cornerPoints := [][]float64{{}, {}}
	for _, p := range shellCornedPoints {
		cornerPoints[0] = append(cornerPoints[0], p.X)
		cornerPoints[1] = append(cornerPoints[1], p.Y)
	}
	fmt.Print(cornerPoints)
	plot.AddPointGroup("Corner points", "lines", cornerPoints)

	// Set min and max values for axes
	minX, maxX := MinMax(allPoints[0])
	plot.SetXrange(int(minX)-2, int(maxX)+2)
	minY, maxY := MinMax(allPoints[1])
	plot.SetYrange(int(minY)-2, int(maxY)+2)

	// Save png
	err = plot.SavePlot(output)
	if err != nil {
		return err
	}

	return nil
}

// Compare numbers in array
func MinMax[K constraints.Integer | constraints.Float](arr []K) (K, K) {
	return pie.Min(arr), pie.Max(arr)
}

// Find most left on X axe point
func LeftXPoint(points []Point) Point {
	xPoints := pie.Map(points, func(p Point) float64 { return p.X })
	min, _ := MinMax(xPoints)

	return pie.Filter(points, func(p Point) bool {
		return p.X == min
	})[0]
}

// Find most right on X axe point
func RightXPoint(points []Point) Point {
	xPoints := pie.Map(points, func(p Point) float64 { return p.X })
	_, max := MinMax(xPoints)

	return pie.Filter(points, func(p Point) bool {
		return p.X == max
	})[0]
}

type Position int32

const (
	OnLine Position = iota
	Right
	Left Position = -1
)

func (p *Point) SideByLine(startPoint, endPoint Point) Position {
	position := (endPoint.X-startPoint.X)*(p.Y-startPoint.Y) - (endPoint.Y-startPoint.Y)*(p.X-startPoint.X)
	if position > 0 {
		return Left
	} else if position < 0 {
		return Right
	} else {
		return OnLine
	}
}

func pointsAtLeftSide(a, b Point, points []Point) []Point {
	return pie.Filter(points, func(p Point) bool {
		return p.SideByLine(a, b) == Left
	})
}

func QuickHull(points []Point) []Point {
	// Left x point
	maxLeft := LeftXPoint(points)
	// Right x point
	maxRight := RightXPoint(points)
	// Points at left side
	s1 := pointsAtLeftSide(maxLeft, maxRight, points)
	// Points at right side
	s2 := pointsAtLeftSide(maxRight, maxLeft, points)

	leftHull := QuickHullHelper(maxLeft, maxRight, s1)
	rightHull := QuickHullHelper(maxLeft, maxRight, s2)

	res := []Point{maxLeft, maxRight}
	res = append(res, rightHull...)
	res = append(res, leftHull...)
	return pie.Unique(res)
}

func QuickHullHelper(a, b Point, points []Point) []Point {

}
