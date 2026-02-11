package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Objetivo: Implementar Productor–Consumidor con canales.
// Un productor genera N valores y los envía por un canal; varios consumidores los procesan.
// Practicar cierre de canal y uso de WaitGroup.
// TODO: completa los pasos marcados.

func productor(n int, out chan<- int) {
	defer close(out) // TODO: ✓
	for i := 1; i <= n; i++ {
		v := rand.Intn(100)
		fmt.Printf("[productor] envía %d\n", v)
		out <- v
		// TODO: ✓
		ms := 100 + rand.Intn(299)
        time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

func consumidor(id int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range in { // TODO:✓
		fmt.Printf("[consumidor %d] recibe %d\n", id, v)
		// TODO: ✓
		ms := 100 + rand.Intn(200)
        time.Sleep(time.Duration(ms) * time.Millisecond)
	}
	fmt.Printf("[consumidor %d] canal cerrado, termina\n", id)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	valores := 10
	consumidores := 3

	ch := make(chan int, 4) // TODO: ✓
    
	var wg sync.WaitGroup
	wg.Add(consumidores)
	// TODO: ✓
	for i := 1; i <= consumidores; i++ {
    go consumidor(i, ch, &wg)
	}

	go productor(valores, ch)

	wg.Wait()
	fmt.Println("Listo: todos los consumidores terminaron.")
}
