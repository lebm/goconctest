package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	logch := make(chan string)
	// ch used by log to signal that it ended processing messages
	logend := make(chan bool)

	go loga(logch, logend)

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		id := i

		go func() {
			defer wg.Done()
			worker(id, logch)
		}()
	}

	wg.Wait()
	// closing logch signals loga that there is no more message to process
	close(logch)
	// ...but we have to wait loga signal that it ended processing messages. It uses logend ch. to signal it.
	// If the progam ends before confirmation, we can loose the last message.
	<-logend

}

func worker(id int, ch chan string) {
	ch <- fmt.Sprintf("Worker %d starting", id)
	time.Sleep(time.Second * time.Duration(rand.Intn(3)))
	ch <- fmt.Sprintf("Worker %d done", id)
}

func loga(ch chan string, finish chan bool) {
	for v := range ch {
		fmt.Println(v)
	}
	finish <- true
}
