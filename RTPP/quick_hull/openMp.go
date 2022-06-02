package quick_hull

import (
	"sync"

	"github.com/elliotchance/pie/v2"
)

type Result struct {
	sync.Mutex
	points []Point
}

// First step of algorithm
func QuickHullOpenMp(points []Point) []Point {
	// Get local axis max points
	maxLeft, maxRight := getMaxPoints(points)

	res := &Result{points: []Point{}}
	res.Lock()
	res.points = append(res.points, maxLeft)
	res.points = append(res.points, maxRight)
	res.Unlock()

	// Run in two goroutines
	// Points at left side
	s1, s2 := pointsAtReverseSides(maxLeft, maxRight, points)

	var wg sync.WaitGroup
	// Run recursive steps
	wg.Add(2)
	go func() {
		defer wg.Done()
		QuickHullOpenMpHelper(res, maxLeft, maxRight, s1)
	}()
	go func() {
		defer wg.Done()
		QuickHullOpenMpHelper(res, maxRight, maxLeft, s2)
	}()
	wg.Wait()

	res.Lock()
	// Form result
	result := pie.SortStableUsing(pie.Unique(res.points), func(a, b Point) bool {
		if a.Name < b.Name {
			return true
		} else {
			return false
		}
	})
	res.Unlock()
	return result
}

// Second and more recursive steps of an algorithm
func QuickHullOpenMpHelper(res *Result, a, b Point, points []Point) {
	var wg sync.WaitGroup
	// // Case than triangle variant is one
	if len(points) <= 1 {
		res.Lock()
		res.points = append(res.points, points...)
		res.Unlock()
		return
	}

	// Most distant point from ab
	h := MostDistantPointToLine(a, b, points)
	s1, s2 := pointsAtSides(a, h, b, points)
	res.Lock()
	res.points = append(res.points, h)
	res.Unlock()

	// Run recursive steps
	wg.Add(2)
	go func() {
		defer wg.Done()
		QuickHullOpenMpHelper(res, a, h, s1)
	}()
	go func() {
		defer wg.Done()
		QuickHullOpenMpHelper(res, h, b, s2)
	}()
}

func getMaxPoints(points []Point) (Point, Point) {
	var wg sync.WaitGroup

	var maxPoints struct {
		sync.Mutex
		maxLeft  Point
		maxRight Point
	}

	wg.Add(2)
	// Left x point
	go func() {
		defer wg.Done()
		left := LeftXPoint(points)
		maxPoints.Lock()
		maxPoints.maxLeft = left
		maxPoints.Unlock()
	}()
	// Right x point
	go func() {
		defer wg.Done()
		right := RightXPoint(points)
		maxPoints.Lock()
		maxPoints.maxRight = right
		maxPoints.Unlock()
	}()
	wg.Wait()
	return maxPoints.maxLeft, maxPoints.maxRight
}

func pointsAtSides(a, h, b Point, points []Point) ([]Point, []Point) {
	var wg sync.WaitGroup
	var sidePoints struct {
		sync.Mutex
		leftSide  []Point
		rightSide []Point
	}

	wg.Add(2)
	// Run in two goroutines
	// Points at left side
	go func() {
		defer wg.Done()
		leftSide := PointsAtLeftSide(a, h, points)
		sidePoints.Lock()
		sidePoints.leftSide = leftSide
		sidePoints.Unlock()
	}()
	// Points at right side
	go func() {
		defer wg.Done()
		rightSide := PointsAtLeftSide(h, b, points)
		sidePoints.Lock()
		sidePoints.rightSide = rightSide
		sidePoints.Unlock()
	}()
	wg.Wait()

	return sidePoints.leftSide, sidePoints.rightSide

}

func pointsAtReverseSides(a, b Point, points []Point) ([]Point, []Point) {
	var wg sync.WaitGroup
	var sidePoints struct {
		sync.Mutex
		leftSide  []Point
		rightSide []Point
	}

	wg.Add(2)

	// Run in two goroutines
	// Points at left side
	go func() {
		defer wg.Done()
		leftSide := PointsAtLeftSide(a, b, points)
		sidePoints.Lock()
		sidePoints.leftSide = leftSide
		sidePoints.Unlock()
	}()
	// Points at right side
	go func() {
		defer wg.Done()
		rightSide := PointsAtLeftSide(b, a, points)
		sidePoints.Lock()
		sidePoints.rightSide = rightSide
		sidePoints.Unlock()
	}()
	wg.Wait()

	return sidePoints.leftSide, sidePoints.rightSide

}
