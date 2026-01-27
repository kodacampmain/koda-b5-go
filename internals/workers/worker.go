package workers

import (
	"fmt"
	"sync"
	"time"
)

func WorkerUnion() {
	c := make(chan int)
	var wg sync.WaitGroup

	for i := range 2 {
		wg.Add(1)
		go worker(i, c, &wg)
	}

	for i := range 10 {
		c <- i
	}
	close(c)
	wg.Wait()
}

func worker(id int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		msg, ok := <-c
		if !ok {
			break
		}
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("Worker ID: %d, Message: %d\n", id, msg)
	}
}
