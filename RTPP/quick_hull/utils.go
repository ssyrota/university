package quick_hull

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/elliotchance/pie/v2"
	"golang.org/x/exp/constraints"
)

// Plot point
type Point struct {
	X    float64
	Y    float64
	Name string
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

// Enumerate points by centroid angle
func EnumPossiblePoints(points []Point) []Point {
	centroid := defineCentroid(points)
	fmt.Println(centroid)
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
		rp.p.Name = strconv.FormatInt(int64(i+1), 10)
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

// Compare numbers in array
func MinMax[K constraints.Integer | constraints.Float](arr []K) (K, K) {
	return pie.Min(arr), pie.Max(arr)
}

// Helper for concatenating many lists
func ConcatList[T any](lists ...[]T) []T {
	res := []T{}
	for _, v := range lists {
		res = append(res, v...)
	}
	return res
}
