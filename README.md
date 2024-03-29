![gool](img/gool.png)

Lightweight goroutine pool / 轻量级Golang协程池

## Installation

To install this package, you need to install Go and set your Go workspace first.

1. You first need [Go](https://golang.org/) installed, then you can use the below Go command to install gool.

```sh
go get -u github.com/LgoLgo/gool
```

2. Import it in your code:

```go
import "github.com/LgoLgo/gool"
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

	"github.com/LgoLgo/gool/pool"
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

## License

This project is under the Apache License 2.0. See the LICENSE file for the full license text.