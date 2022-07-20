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
	// ch usado por loga p/ indicar que acabou de processar tudo e saiu do loop
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
	// fecha logch para indicar para loga que não há mais mensagens...
	close(logch)
	// ...mas tem que esperar loga indicar que terminou tudo antes de encerrar o programa.
	// Se não esperar a confirmação de log a última mensagem pode ser perdida.
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
