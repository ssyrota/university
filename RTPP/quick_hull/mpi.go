package quick_hull

import (
	"github.com/elliotchance/pie/v2"
)

// First step of algorithm
func QuickHullMpi(points []Point) []Point {
	// Run in two goroutines
	maxLeftCh := RunInGoroutine(LeftXPoint, points)
	maxRightCh := RunInGoroutine(RightXPoint, points)

	// Left x point
	maxLeft := <-maxLeftCh
	// Right x point
	maxRight := <-maxRightCh

	// Run in two goroutines
	// Points at left side
	s1Chan := make(chan []Point)
	go func() {
		s1Chan <- pointsAtLeftSide(maxLeft, maxRight, points)
	}()
	// Points at right side
	s2Chan := make(chan []Point)
	go func() {
		s2Chan <- pointsAtLeftSide(maxRight, maxLeft, points)
	}()

	// Run recursive steps
	leftHullChan := make(chan []Point)
	go func() {
		leftHullChan <- QuickHullMpiHelper(maxLeft, maxRight, <-s1Chan)
	}()
	rightHullChan := make(chan []Point)
	go func() {
		rightHullChan <- QuickHullMpiHelper(maxRight, maxLeft, <-s2Chan)
	}()

	// Form result
	res := pie.SortStableUsing(pie.Unique(
		ConcatList([]Point{
			maxLeft,
			maxRight},
			<-rightHullChan,
			<-leftHullChan)),
		func(a, b Point) bool {
			if a.Name < b.Name {
				return true
			} else {
				return false
			}
		})

	return res
}

// Second and more recursive steps of an algorithm
func QuickHullMpiHelper(a, b Point, points []Point) []Point {
	// Case than triangle variant is one
	if len(points) <= 1 {
		return points
	}

	// Most distant point from ab
	h := MostDistantPointToLine(a, b, points)

	// Run in two goroutines
	// Points at left side
	s1Chan := make(chan []Point)
	go func() {
		s1Chan <- pointsAtLeftSide(a, h, points)
	}()
	// Points at right side
	s2Chan := make(chan []Point)
	go func() {
		s2Chan <- pointsAtLeftSide(h, b, points)
	}()

	// Run recursive steps
	leftHullChan := make(chan []Point)
	go func() {
		leftHullChan <- QuickHullMpiHelper(a, h, <-s1Chan)
	}()
	rightHullChan := make(chan []Point)
	go func() {
		rightHullChan <- QuickHullMpiHelper(h, b, <-s2Chan)
	}()

	res := ConcatList(<-rightHullChan, <-leftHullChan, []Point{h})
	return res
}

func RunInGoroutine[T any, U any](f func(U) T, args U) chan T {
	ch := make(chan T)
	go func() {
		ch <- f(args)
	}()
	return ch
}
