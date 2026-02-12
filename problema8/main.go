package main

import (
	"fmt"
	"time"
	"sync" //agregue
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// TODO: completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		// TODO: simular trabajo
        time.Sleep(1 * time.Second)
		ch <- x * x
	}()
	return ch
}

func fanIn(canales ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(canales))
	for _, c := range canales {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
func main() {
	fmt.Println("cálculos")
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(5)
	f3 := asyncCuadrado(8)
	resultadoCombinado := fanIn(f1, f2, f3)
	for res := range resultadoCombinado {
		fmt.Printf("Resultado: %d\n", res)
	}
	fmt.Println("terminado.")
}