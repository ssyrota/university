package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

type Philosopher struct {
	id   int
	eats int
}

// Lock forks with some delay for proof of deadlock free
func (p *Philosopher) takeForks() {
	p.folkAction()
}

// Unlock forks with some delay
func (p *Philosopher) leftForks() {
	p.folkAction()
}

// Simulate time operation
func (p *Philosopher) eatAction() {
	fmt.Printf("[%d] starting to eat\n", p.id)
	p.folkAction()
	fmt.Printf("[%d] finishing eating\n", p.id)
	p.eats++
}

// Random sleep
func (p *Philosopher) folkAction() {
	time.Sleep(time.Second)
}

// Run eating action
func (p Philosopher) eat() {
	for {
		requestEat <- p.id
		if eat := <-startEat; eat == 1 {
			p.takeForks()
			p.eatAction()
			finishEat <- p.id
			p.leftForks()
		}
	}
}

// Wg for each philossoper
var wgEat sync.WaitGroup

// Philossopers amount
var count = 5

// Waiter chanels
var finishEat = make(chan int)
var startEat = make(chan int)
var requestEat = make(chan int)

type Statuses struct {
	sync.Mutex
	data map[int]bool
}

var eatingStatuses = &Statuses{data: make(map[int]bool)}

func waiter() {
	for {
		select {
		case id := <-requestEat:
			eatingStatuses.Lock()

			leftId := id - 1
			if leftId == 0 {
				leftId = count
			}
			rightId := id + 1
			if rightId > count {
				rightId = 1
			}

			if eatingStatuses.data[leftId] || eatingStatuses.data[rightId] {
				eatingStatuses.Unlock()
				startEat <- 0
			} else {
				eatingStatuses.data[id] = true
				eatingStatuses.Unlock()
				startEat <- 1
			}

		case id := <-finishEat:
			eatingStatuses.Lock()
			eatingStatuses.data[id] = false
			eatingStatuses.Unlock()
		}
	}
}

func main() {
	philossophers := make([]*Philosopher, count)
	for i := 0; i < count; i++ {
		philossophers[i] = &Philosopher{
			id:   i + 1,
			eats: 0,
		}
	}
	go waiter()
	for _, v := range philossophers {
		wgEat.Add(1)
		go v.eat()
	}
	wgEat.Wait()
}
