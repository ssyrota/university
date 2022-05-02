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
		if p.eats == 3 {
			finishEat <- p.id
			wgEat.Done()
			fmt.Printf("[%d] left table\n", p.id)
			return
		}

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
var count = 10

// Waiter chanels
var finishEat = make(chan int)
var startEat = make(chan int)
var requestEat = make(chan int)

func waiter() {
	eatingStatuses := make(map[int]bool)
	for {
		select {
		case id := <-requestEat:
			leftId := id - 1
			if leftId == 0 {
				leftId = count
			}
			rightId := id + 1
			if rightId > count {
				rightId = 1
			}

			if eatingStatuses[leftId] || eatingStatuses[rightId] {
				startEat <- 0
			} else {
				eatingStatuses[id] = true
				startEat <- 1
			}

		case id := <-finishEat:
			eatingStatuses[id] = false
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
