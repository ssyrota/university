package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/elliotchance/pie/v2"
	mpi "github.com/sbromberger/gompi"
	qh "github.com/sergeycrisp/university/RTPP/quick_hull"
)

const MasterRank = 0

func main() {
	mpi.Start(true)
	worldRank := mpi.WorldRank()
	worldComm := mpi.NewCommunicator(nil)
	worldSize := worldComm.Size()

	if worldSize == 1 {
		panic("1 proc is to small capacity for start")
	}

	if worldRank == MasterRank {
		mainProcess(worldComm)
	} else {
		subProcess(worldComm)
	}

	mpi.Stop()

}

// Main process
func mainProcess(worldComm *mpi.Communicator) {
	worldSize := worldComm.Size()
	fmt.Printf("Processors: %d\n", worldSize)
	// Form sorted points
	points := qh.EnumPossiblePoints(GeneratePoints(5000))

	start := time.Now()

	// Left x point
	maxLeft := qh.LeftXPoint(points)
	// Right x point
	maxRight := qh.RightXPoint(points)

	// Points at left side
	s1 := qh.PointsAtLeftSide(maxLeft, maxRight, points)
	// Points at right side
	s2 := qh.PointsAtLeftSide(maxRight, maxLeft, points)

	// Track used processes
	usedProcesses, performedPoints := distributedQuickHull(maxLeft, maxRight, s1, s2, worldComm)

	var resultsSync []qh.Point
	if worldSize == 2 {
		pts1 := qh.ReceivePoints(worldComm, usedProcesses[0], qh.MainAlgorithmCh)
		resultsSync = append(resultsSync, pts1...)
		pts2 := qh.ReceivePoints(worldComm, usedProcesses[0], qh.MainAlgorithmCh)
		resultsSync = append(resultsSync, pts2...)
	} else {
		for _, pId := range usedProcesses {
			pts := qh.ReceivePoints(worldComm, pId, qh.MainAlgorithmCh)
			resultsSync = append(resultsSync, pts...)
		}
	}
	// Form result
	pie.SortStableUsing(
		pie.Unique(
			qh.ConcatList(
				append([]qh.Point{maxLeft, maxRight}, performedPoints...),
				resultsSync)),
		qh.CompareByName)

	elapsed := time.Since(start)
	fmt.Printf("Performed time: %v ms\n", elapsed.Milliseconds())
}

func distributedQuickHull(maxLeft, maxRight qh.Point, s1, s2 []qh.Point, worldComm *mpi.Communicator) ([]int, []qh.Point) {
	// Count processes
	worldSize := worldComm.Size()
	var processes []int
	for pId := 1; pId < worldSize; pId++ {
		processes = append(processes, pId)
	}
	// 1 worker
	if len(processes) == 1 {
		qh.SendPoints(qh.ConcatList([]qh.Point{maxLeft, maxRight}, s1), worldComm, processes[0], qh.MainAlgorithmCh)
		qh.SendPoints(qh.ConcatList([]qh.Point{maxRight, maxLeft}, s2), worldComm, processes[0], qh.MainAlgorithmCh)
		return processes, []qh.Point{}
	}

	// 2 workers
	if len(processes) == 2 {
		qh.SendPoints(qh.ConcatList([]qh.Point{maxLeft, maxRight}, s1), worldComm, processes[0], qh.MainAlgorithmCh)
		qh.SendPoints(qh.ConcatList([]qh.Point{maxRight, maxLeft}, s2), worldComm, processes[1], qh.MainAlgorithmCh)
		return processes, []qh.Point{}
	}

	// 3 workers
	if len(processes) == 3 {
		h, s1part1, s1part2 := qh.QuickHullSeparator(maxLeft, maxRight, s1)
		qh.SendPoints(qh.ConcatList([]qh.Point{maxLeft, h}, s1part1), worldComm, processes[0], qh.MainAlgorithmCh)
		qh.SendPoints(qh.ConcatList([]qh.Point{h, maxRight}, s1part2), worldComm, processes[1], qh.MainAlgorithmCh)
		qh.SendPoints(qh.ConcatList([]qh.Point{maxLeft, maxRight}, s2), worldComm, processes[2], qh.MainAlgorithmCh)
		return processes, []qh.Point{h}
	}

	panic("not-implemented-scalability")

}

// Second and more recursive steps of an algorithm
func QuickHullHelperMpi(a, b qh.Point, points []qh.Point) []qh.Point {
	// Case than triangle variant is one
	if len(points) <= 1 {
		return points
	}

	h, s1, s2 := qh.QuickHullSeparator(a, b, points)

	// Run recursive steps
	leftHull := QuickHullHelperMpi(a, h, s1)
	rightHull := QuickHullHelperMpi(h, b, s2)

	res := qh.ConcatList(rightHull, leftHull, []qh.Point{h})
	return res
}

// Sub process
func subProcess(worldComm *mpi.Communicator) {
	if worldComm.Size() == 2 {
		receiveAndPerformAlgorithm(worldComm)
		receiveAndPerformAlgorithm(worldComm)
	} else {
		receiveAndPerformAlgorithm(worldComm)
	}

}

func receiveAndPerformAlgorithm(worldComm *mpi.Communicator) {
	// Receive points to work with main algorithm
	points := qh.ReceivePoints(worldComm, MasterRank, qh.MainAlgorithmCh)
	a := pie.Pop(&points)
	b := pie.Pop(&points)
	res := QuickHullHelperMpi(*a, *b, points)
	qh.SendPoints(res, worldComm, MasterRank, qh.MainAlgorithmCh)
}

//Generate points
func GeneratePoints(n int32) []qh.Point {
	res := make([]qh.Point, n)
	var i int32
	for i = 0; i < n; i++ {
		res[i] = qh.Point{X: rand.Float64() * float64(n), Y: rand.Float64() * float64(n), Name: i}
	}
	return res
}
