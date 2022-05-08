package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	input := 1
	fChan := make(chan bool)
	gChan := make(chan bool)
	ctx, cancel := context.WithCancel(context.Background())
	results := make(map[string]bool)

	go runFunctionInThread(ctx, input, fChan, f)
	go runFunctionInThread(ctx, input, gChan, g)
	ticker := time.NewTicker(time.Second)

waiter:
	for {
		select {
		case <-ticker.C:
			ticker.Stop()
			fmt.Print(alert)
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			answer := input.Text()

			if answer == "0" {
				ticker.Reset(time.Second)
				continue
			}

			if answer == "1" {
				cancel()
				fmt.Println("[exit] program ended")
				break waiter
			}

			if answer == "2" {
				ticker.Stop()
				continue
			}

		case resF := <-fChan:
			if resF {
				fmt.Println(true)
				break waiter
			}
			if _, ok := results["G"]; ok {
				fmt.Print(false)
				break waiter
			}
			results["F"] = resF

		case resG := <-gChan:
			if resG {
				fmt.Println(true)
				break waiter
			}
			if _, ok := results["F"]; ok {
				fmt.Print(false)
				break waiter
			}
			results["G"] = resG
		}

	}

}

func runFunctionInThread(ctx context.Context, input int, chanFunction chan bool, function func(a int) bool) {
	internalCh := make(chan bool)
	go func() {
		internalCh <- function(input)
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case val := <-internalCh:
			chanFunction <- val
			return
		}
	}
}

func f(x int) bool {
	fmt.Println("[started]-f")
	time.Sleep(time.Second * 2)
	return true
}

func g(x int) bool {
	fmt.Println("[started]-g")
	for {
		x = x + x
	}
}

var alert = `
Program execution time is long. 
continue (0), stop(1) or continue without ask again(2)?
Enter: `
