package quick_hull

import (
	"github.com/elliotchance/pie/v2"
)

// First step of algorithm
func QuickHullSequential(points []Point) []Point {
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
	rightHull := QuickHullHelper(maxRight, maxLeft, s2)

	// Form result
	res := pie.SortStableUsing(pie.Unique(ConcatList([]Point{maxLeft, maxRight}, rightHull, leftHull)), func(a, b Point) bool {
		if a.Name < b.Name {
			return true
		} else {
			return false
		}
	})

	return res
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

	res := ConcatList(rightHull, leftHull, []Point{h})
	return res
}
