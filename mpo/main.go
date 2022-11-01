package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var timeout = time.Second * 1

func main() {
	mod := os.Args[1]
	if mod != "right" && mod != "wrong" {
		log.Fatalf("error: %s", "runner type should be passed via args")
	}
	N := uint(10)
	K := uint(10)
	P := 0.01 /* 1% */

	if mod == "right" {
		safeSimulator := NewCrystal[CrystalSafe](N, K, P)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			safeSimulator.BrownianMotionSimulate(timeout)
		}()
		wg.Wait()
		safeSimulator.PrintTotalAtoms("safe")
	} else {
		unsafeSimulator := NewCrystal[CrystalUnsafe](N, K, P)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			unsafeSimulator.BrownianMotionSimulate(timeout)
		}()
		wg.Wait()
	}
	/*
		В кінці програми необхідне переобчислити загальну кількість атомів домішки (вона не повинна змінитися) і вивести його на екран.
	*/
}

/*
Нехай є простір, що є кінцевим набором кліток N, по яких розподілене деяка кількість K атомів домішки.
У кожен момент часу атом може перейти в будь-яку з двох сусідніх кліток з однією і тією ж вірогідністю р
*/
func NewCrystal[T CrystalUnsafe | CrystalSafe](n uint, k uint, p float64) *T {
	cells := make([]int32, n)

	// У початковому стані всі атоми домішки зосередженні зліва.
	cells[0] = int32(k)
	return &T{
		Crystal{
			cells: cells,
			p:     p,
		},
	}
}

type BrownianMotionSimulator interface {
	BrownianMotionSimulate(timeout time.Duration)
}

var _ BrownianMotionSimulator = new(CrystalSafe)
var _ BrownianMotionSimulator = new(CrystalUnsafe)

type CrystalSafe struct {
	Crystal
}

// BrownianMotionSimulator implementation.
func (cs *CrystalSafe) BrownianMotionSimulate(timeout time.Duration) {
	maxCellIdx := len(cs.cells) - 1
	defaultAtomIndex := 0
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printSnapshotWithIvl(cs.cells, ctx)
	}()
	for i := int32(0); i < cs.cells[0]; i++ {
		wg.Add(1)
		go func() {
			curIdx := defaultAtomIndex
			defer wg.Done()
		swapping:
			for {
				if !shouldMove(cs.p) {
					continue swapping
				}
				moveSide := right
				if rightLeft() {
					moveSide = left
				}
				if curIdx == defaultAtomIndex {
					moveSide = right
				} else if curIdx == maxCellIdx {
					moveSide = left
				}

				if moveSide == right {
					atomic.AddInt32(&cs.cells[curIdx], -1)
					curIdx += 1
					atomic.AddInt32(&cs.cells[curIdx], 1)
				} else {
					atomic.AddInt32(&cs.cells[curIdx], -1)
					curIdx -= 1
					atomic.AddInt32(&cs.cells[curIdx], 1)

				}
				select {
				case <-ctx.Done():
					return
				default:
					continue swapping
				}
			}
		}()
	}
	wg.Wait()
}

type CrystalUnsafe struct {
	Crystal
}

// BrownianMotionSimulator implementation.
func (cs *CrystalUnsafe) BrownianMotionSimulate(timeout time.Duration) {
	maxCellIdx := len(cs.cells) - 1
	defaultAtomIndex := 0
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printSnapshotWithIvl(cs.cells, ctx)
	}()
	for i := int32(0); i < cs.cells[0]; i++ {
		wg.Add(1)
		go func() {
			curIdx := defaultAtomIndex
			defer wg.Done()
		swapping:
			for {
				if !shouldMove(cs.p) {
					continue swapping
				}
				moveSide := right
				if rightLeft() {
					moveSide = left
				}
				if curIdx == defaultAtomIndex {
					moveSide = right
				} else if curIdx == maxCellIdx {
					moveSide = left
				}

				if moveSide == right {
					tmpVal := cs.cells[curIdx]
					cs.cells[curIdx] = tmpVal - 1
					curIdx += 1
					tmpVal = cs.cells[curIdx]
					cs.cells[curIdx] = tmpVal + 1
				} else {
					tmpVal := cs.cells[curIdx]
					cs.cells[curIdx] = tmpVal - 1
					curIdx -= 1
					tmpVal = cs.cells[curIdx]
					cs.cells[curIdx] = tmpVal + 1

				}
				select {
				case <-ctx.Done():
					return
				default:
					continue swapping
				}
			}
		}()
	}
	wg.Wait()
}

type Crystal struct {
	cells []int32
	p     float64
}

func (c *Crystal) PrintTotalAtoms(prefix string) {
	sum := int32(0)
	for _, v := range c.cells {
		sum += v
	}
	fmt.Printf("%s: %v\n", prefix, sum)
}

func shouldMove(p float64) bool {
	return rand.Float64() < p
}

func rightLeft() bool {
	return rand.Float64() < 0.5
}

type Side bool

const (
	right Side = true
	left  Side = false
)

var snapShotInterval = time.Millisecond * 100

func printSnapshotWithIvl(state []int32, ctx context.Context) {
	ticker := time.NewTicker(snapShotInterval)
	for {
		select {
		case <-ctx.Done():
			return
		default:
			<-ticker.C
			sum := int32(0)
			for _, v := range state {
				sum += v
			}
			fmt.Printf("Sum: %v\n", sum)
			log.Printf("%v", state)
		}
	}
}
