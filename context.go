package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func contextDemo() {
	ContextWithCancel()
	ContextWithTimeout()
	// TODO: Implement ContextWithValue()
}

func ContextWithCancel() {
	var wg sync.WaitGroup
	wg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()
		getRunningContainers(ctx)
	}()

	go func() {
		defer wg.Done()
		createContainer(ctx, "ubuntu")
	}()

	<-time.After(time.Second * 3)
	cancel()
}

func ContextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		getRunningContainers(ctx)
	}()

	go func() {
		defer wg.Done()
		createContainer(ctx, "ubuntu")
	}()

	wg.Wait()
}

func getRunningContainers(ctx context.Context) {
	select {
	case <-ctx.Done():
		log.Println("getRunningContainers: timeout")
		return
	case <-time.After(time.Second * 2):
		log.Println("getRunningContainers: process completed")
	}
}

func createContainer(ctx context.Context, name string) {
	select {
	case <-ctx.Done():
		log.Println("createContainer: timeout")
		return
	case <-time.After(time.Second * 2):
		log.Printf("createContainer: process completed %v", name)
	}
}
