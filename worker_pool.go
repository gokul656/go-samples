package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	wg         *sync.WaitGroup
	maxWorkers int      // number of cores 4
	jobs       chan Job // setTimeout, setInterval, fetch

	idleTimeout      time.Duration
	waitingQueueSize int
}

type Job struct {
	f func(n int)
}

func (wp *WorkerPool) AddJob(j Job) {
	wp.jobs <- j
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.maxWorkers; i++ {
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(n int) {
	for job := range wp.jobs {
		wp.wg.Add(1)
		job.f(n)
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Shutdown() {
	close(wp.jobs)
	wp.wg.Wait()
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		maxWorkers:       workers,
		wg:               &sync.WaitGroup{},
		jobs:             make(chan Job, workers),
		idleTimeout:      time.Second * 5,
		waitingQueueSize: 64,
	}
}

func workerPoolExample() {
	pool := NewWorkerPool(4)
	pool.Start()

	for i := 0; i <= 10; i++ {
		pool.AddJob(Job{f: doSomething})
	}

	pool.Shutdown()
}

func doSomething(n int) {
	fmt.Printf("Processing on thread %d\n", n)
	time.Sleep(time.Second * 2)
}
