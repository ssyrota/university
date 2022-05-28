package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/Arafatk/glot"
	"github.com/elliotchance/pie/v2"
	"golang.org/x/exp/constraints"
)

func main() {
	all := []Point{{1, 2, "1"}, {2, 1, "2"}, {4, 3, "3"}, {5, 3, "4"}, {6, 1, "5"}, {7, 4, "6"}, {8, 2, "7"}, {9, 3, "8"}, {10, 3, "9"}}

	// MakePlot(all, all, "2.png")
	fmt.Print(EnumPossiblePoints(all))
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

// Available positions for the point related to the line
type Position int32

const (
	OnLine Position = iota
	Right
	Left Position = -1
)

// Define a position of point related to the line
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

// Filter points only from left side of the line
func pointsAtLeftSide(a, b Point, points []Point) []Point {
	return pie.Filter(points, func(p Point) bool {
		return p.SideByLine(a, b) == Left
	})
}

// First step of algorithm
func QuickHull(points []Point) []Point {
	// Left x point
	maxLeft := LeftXPoint(points)
	// Right x point
	maxRight := RightXPoint(points)
	// Points at left side
	s1 := pointsAtLeftSide(maxLeft, maxRight, points)
	// Points at right side
	s2 := pointsAtLeftSide(maxRight, maxLeft, points)

	// Run recursive steps
	leftHull := QuickHullHelper(maxLeft, maxRight, s1)
	rightHull := QuickHullHelper(maxLeft, maxRight, s2)

	// Form result
	return pie.Unique(concatList([]Point{maxLeft, maxRight}, rightHull, leftHull))
}

// Second and more recursive steps of an algorithm
func QuickHullHelper(a, b Point, points []Point) []Point {
	// Case than triangle variant is one
	if len(points) <= 1 {
		return points
	}

	// Most distant point from ab
	h := MostDistantPointToLine(a, b, points)
	// Points at left side
	s1 := pointsAtLeftSide(a, h, points)
	// Points at right side
	s2 := pointsAtLeftSide(h, b, points)

	// Run recursive steps
	leftHull := QuickHullHelper(a, h, s1)
	rightHull := QuickHullHelper(h, b, s2)

	return concatList(rightHull, leftHull, []Point{h})
}

// Find most distant point related to line
func MostDistantPointToLine(a, b Point, points []Point) Point {
	return pie.SortStableUsing(points, func(prev, next Point) bool {
		if TriangleShape(a, b, prev) > TriangleShape(a, b, next) {
			return true
		} else {
			return false
		}
	})[0]
}

// Find triangle shape
func TriangleShape(a, b, c Point) float64 {
	return (1 / 2.0) * math.Abs((b.X-a.X)*(c.Y-a.Y)-(c.X-a.X)*(b.Y-a.Y))
}

// Helper for concatenating many lists
func concatList[T any](lists ...[]T) []T {
	res := []T{}
	for _, v := range lists {
		res = append(res, v...)
	}
	return res
}

// Enumerate points by centroid angle
func EnumPossiblePoints(points []Point) []Point {
	centroid := defineCentroid(points)
	type RelativePosition struct {
		p     *Point
		angle float64
	}
	res := pie.SortStableUsing(pie.Map(points, func(p Point) RelativePosition {
		return RelativePosition{&p, RadiansFromCentroid(centroid, p)}
	}), func(a, b RelativePosition) bool {
		if a.angle < b.angle {
			return true
		} else {
			return false
		}
	})

	for i, rp := range res {
		rp.p.Name = strconv.FormatInt(int64(i), 10)
	}
	return pie.Map(res, func(rp RelativePosition) Point {
		return *rp.p
	})
}

// Radians from centroid to point
func RadiansFromCentroid(centroid, p Point) float64 {
	return math.Atan2(p.Y-centroid.Y, p.X-centroid.X)
}

// Get centroid from all points. Named q
func defineCentroid(points []Point) Point {
	shuffled := pie.Shuffle(points, rand.New(rand.NewSource(time.Now().UnixNano())))
	a := pie.Pop(&shuffled)
	b := pie.Pop(&shuffled)
	c := pie.Pop(&shuffled)

	return Point{(a.X + b.X + c.X) / 3, (a.Y + b.Y + c.Y) / 3, "centroid"}
}
