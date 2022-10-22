package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	timeout := time.Second * 10
	N := uint(1000000) /* 1M */
	K := uint(10000)   /* 10k*/
	P := 0.01          /* 1% */

	safeSimulator := NewCrystal[CrystalSafe](N, K, P)
	unsafeSimulator := NewCrystal[CrystalUnsafe](N, K, P)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeSimulator.BrownianMotionSimulate(timeout)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		unsafeSimulator.BrownianMotionSimulate(timeout)
	}()
	wg.Wait()
	/*
		В кінці програми необхідне переобчислити загальну кількість атомів домішки (вона не повинна змінитися) і вивести його на екран.
	*/
	safeSimulator.PrintTotalAtoms("safe")
	unsafeSimulator.PrintTotalAtoms("unsafe")

}

/*
Нехай є простір, що є кінцевим набором кліток N, по яких розподілене деяка кількість K атомів домішки.
У кожен момент часу атом може перейти в будь-яку з двох сусідніх кліток з однією і тією ж вірогідністю р
*/
func NewCrystal[T CrystalUnsafe | CrystalSafe](n uint, k uint, p float64) *T {
	cells := make([]uint64, n)

	// У початковому стані всі атоми домішки зосередженні зліва.
	cells[0] = uint64(k)
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
	for i := uint64(0); i < cs.cells[0]; i++ {
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
					atomic.AddUint64(&cs.cells[curIdx], ^uint64(0))
					curIdx += 1
					atomic.AddUint64(&cs.cells[curIdx], 1)
				} else {
					atomic.AddUint64(&cs.cells[curIdx], ^uint64(0))
					curIdx -= 1
					atomic.AddUint64(&cs.cells[curIdx], 1)

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
	for i := uint64(0); i < cs.cells[0]; i++ {
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
	cells []uint64
	p     float64
}

func (c *Crystal) PrintTotalAtoms(prefix string) {
	sum := uint64(0)
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
