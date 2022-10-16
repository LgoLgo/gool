package main

import (
	"fmt"
	"time"

	"github.com/LgoLgo/LgoPool/pool"
)

func main() {
	// Create some tasks
	t := pool.NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	// Create a goroutine pool,
	// the maximum number of workers in this goroutine pool is 4
	p := pool.NewPool(4)

	taskNum := 0

	// Give these tasks to the goroutine pool
	go func() {
		for {
			// Continuously write task t to p,
			// each task is to print the current time
			p.EntryChannel <- t
			taskNum += 1
			fmt.Printf("%d tasks currently executed\n", taskNum)
		}
	}()

	// Start the pool and let the pool start working.
	// At this time, the pool will create a worker and let the worker work
	p.Run()
}
