package quick_hull

import (
	"fmt"

	"github.com/elliotchance/pie/v2"
	mpi "github.com/sbromberger/gompi"
)

type PointMI struct {
	x    int
	y    int
	name int
}

const (
	XChanel = iota
	YChanel
	NameChanel
	MaxLeftXChanel
	MaxLeftYChanel
	MaxLeftNameChanel
	MaxRightXChanel
	MaxRightYChanel
	MaxRightNameChanel
	LeftPointsChX
	LeftPointsChY
	LeftPointsChName
	RightPointsChX
	RightPointsChY
	RightPointsChName
)

var MainAlgorithmCh = &PointMI{XChanel, YChanel, NameChanel}
var MaxRightCh = &PointMI{MaxRightXChanel,
	MaxRightYChanel,
	MaxRightNameChanel}
var MaxLeftCh = &PointMI{MaxLeftXChanel,
	MaxLeftYChanel,
	MaxLeftNameChanel}
var LeftPointsCh = &PointMI{LeftPointsChX, LeftPointsChY, LeftPointsChName}
var RightPointsCh = &PointMI{RightPointsChX, RightPointsChY, RightPointsChName}

func QuickHullSeparator(a, b Point, points []Point) (Point, []Point, []Point) {
	// Most distant point from ab
	h := MostDistantPointToLine(a, b, points)
	// Points at left side
	s1 := PointsAtLeftSide(a, h, points)
	// Points at right side
	s2 := PointsAtLeftSide(h, b, points)
	return h, s1, s2
}

// Receive points by MPI chanel
func ReceivePoints(worldComm *mpi.Communicator, from int, pointCh *PointMI) []Point {
	dataX, status := worldComm.RecvFloat64s(from, pointCh.x)
	if status.GetError() != 0 {
		fmt.Println(status.GetError())
	}
	dataY, status := worldComm.RecvFloat64s(from, pointCh.y)
	if status.GetError() != 0 {
		fmt.Println(status.GetError())
	}
	names, status := worldComm.RecvInt32s(from, pointCh.name)
	if status.GetError() != 0 {
		fmt.Println(status.GetError())
	}

	var points []Point
	for i, name := range names {
		points = append(points, Point{Name: name, X: dataX[i], Y: dataY[i]})
	}
	return points
}

// Send points by MPI chanel
func SendPoints(points []Point, worldComm *mpi.Communicator, to int, pointCh *PointMI) {
	dataX := pie.Map(points, func(p Point) float64 { return p.X })
	dataY := pie.Map(points, func(p Point) float64 { return p.Y })
	names := pie.Map(points, func(p Point) int32 { return p.Name })

	worldComm.SendFloat64s(dataX, to, pointCh.x)
	worldComm.SendFloat64s(dataY, to, pointCh.y)
	worldComm.SendInt32s(names, to, pointCh.name)
}

func CompareByName(a, b Point) bool {
	if a.Name < b.Name {
		return true
	} else {
		return false
	}
}
