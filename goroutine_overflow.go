package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func goRoutineOverflow() {

	count := 100

	rountineCount := os.Getenv("ROUTINE_COUNT")
	if rountineCount != "" {
		n, err := strconv.Atoi(rountineCount)
		if err == nil {
			count = n
		}
	}

	wg := &sync.WaitGroup{}
	for i := range count {
		wg.Add(1)
		go print(i, wg)
	}

	wg.Wait()
}

func print(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("processing", i)
	time.Sleep(time.Second * 50)
}
