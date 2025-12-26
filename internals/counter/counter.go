package counter

import (
	"fmt"
)

var number int = 0

// var mtx sync.Mutex
// var mtx sync.RWMutex

var c chan string = make(chan string)

func RunCounter() {
	// var wg sync.WaitGroup
	// c := make(chan string)
	for range 5 {
		// mtx.Lock()
		go incNumber()
		<-c
		// wg.Go(incNumber)
		// wg.Wait()
		// mtx.RLock()
		go showNumber()
		<-c
		// wg.Go(showNumber)
		// wg.Wait()
	}
	// fmt.Println("Done")
	// wg.Wait()
}

func showNumber() {
	// defer mtx.RUnlock()
	// defer func() { c <- "1" }()
	fmt.Println(number)
	c <- ""
}

func incNumber() {
	// defer mtx.Unlock()
	// defer func() { c <- "1" }()
	number++
	c <- ""
}
