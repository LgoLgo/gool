# LgoPool

Lightweight goroutine pool / 轻量级Golang协程池

## Installation

To install Gin package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed (**version 1.15+ is required**), then you can use the below Go command to install Gin.

```sh
go get -u github.com/L2ncE/LgoPool
```

2. Import it in your code:

```go
import "github.com/L2ncE/LgoPool"
```

## Quick start

```sh
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import (
	"fmt"
	"time"

	"github.com/L2ncE/LgoPool/pool"
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
```