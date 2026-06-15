package main

import (
	"fmt"
	"sync"
)

const (
	NumberOfWorkers = 10
	NumberOfTasks   = 100
)

func worker(jobId int, receive <-chan string, res chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for msg := range receive {
		res <- fmt.Sprintf("Received Job ID:%d and Msg: %s", jobId, msg)
	}
}

func main() {
	ch := make(chan string, NumberOfTasks)
	result := make(chan string, NumberOfTasks)
	var wg sync.WaitGroup
	wg.Add(NumberOfWorkers)
	for i := 0; i < NumberOfWorkers; i++ {
		go worker(i, ch, result, &wg)
	}
	for i := 0; i < NumberOfTasks; i++ {
		ch <- fmt.Sprintf("%d", i+1)
	}
	close(ch)
	go func() {
		wg.Wait()
		close(result)
	}()
	for msg := range result {
		fmt.Println(msg)
	}
}
