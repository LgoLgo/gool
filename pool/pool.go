package pool

import "fmt"

type Task struct {
	// There should be a specific business in a Task,
	// the business name is called f
	f func() error
}

// NewTask create a task
func NewTask(argF func() error) *Task {
	t := Task{
		f: argF,
	}

	return &t
}

// Execute to execute task
func (t *Task) Execute() {
	// Call the method that has been bound in the task
	t.f()
}

// Pool define a pool struct
type Pool struct {
	// External Task Entry
	EntryChannel chan *Task

	// Internal Task queue
	JobsChannel chan *Task

	// The maximum number of workers in the goroutine pool
	workerNum int
}

// NewPool Create a new goroutine pool
func NewPool(cap int) *Pool {
	// Create a pool
	p := Pool{
		EntryChannel: make(chan *Task),
		JobsChannel:  make(chan *Task),
		workerNum:    cap,
	}

	return &p
}

// Worker
//
// The coroutine pool creates a Worker and lets this Worker work
func (p *Pool) worker(workerId int) {
	// Permanently fetch jobs from the JobsChannel
	for task := range p.JobsChannel {
		// task is the task that the current worker gets from the JobsChannel
		// Once the task is fetched, execute the task
		task.Execute()
		fmt.Printf("WorkerID:%d completed a task\n", workerId)
	}
}

// Run is the startup method of the goroutine pool
func (p *Pool) Run() {
	// Create a worker to work based on workerNum
	for i := 0; i < p.workerNum; i++ {
		// Each worker should be a goroutine
		go p.worker(i)
	}

	// Get the task from the EntryChannel,
	// and send the fetched task to the JobsChannel
	for task := range p.EntryChannel {
		// Once a task reads
		p.JobsChannel <- task
	}
}
