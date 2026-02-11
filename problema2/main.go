package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Variante insegura (condición de carrera):
func incrementarInseguro(nGoroutines, nIncrementos int) int64 {
	var contador int64 = 0

	var wg sync.WaitGroup
	wg.Add(nGoroutines)

	for i := 0; i < nGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < nIncrementos; j++ {
				// Operación NO atómica: Lectura + Incremento + Escritura
				contador = contador + 1
			}
		}()
	}

	wg.Wait()
	return contador
}

// Variante con Mutex:
func incrementarConMutex(nGoroutines, nIncrementos int) int64 {
	var contador int64 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex // El cerrojo (Lock)

	wg.Add(nGoroutines)

	for i := 0; i < nGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < nIncrementos; j++ {
				mu.Lock()         // Entra a la sección crítica
				contador = contador + 1
				mu.Unlock()       // Sale de la sección crítica
			}
		}()
	}

	wg.Wait()
	return contador
}

// Variante con atomic:
func incrementarConAtomic(nGoroutines, nIncrementos int) int64 {
	var contador int64 = 0
	var wg sync.WaitGroup

	wg.Add(nGoroutines)

	for i := 0; i < nGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < nIncrementos; j++ {
				// Operación atómica a nivel de CPU
				atomic.AddInt64(&contador, 1)
			}
		}()
	}

	wg.Wait()
	return contador
}

func main() {
	// Aumentamos los valores para que la race condition sea más evidente
	nGoroutines := 10
	nIncrementos := 100_000

	fmt.Printf("Configuración: %d goroutines, %d incrementos cada una\n", nGoroutines, nIncrementos)
	fmt.Printf("Esperado correcto: %d\n\n", int64(nGoroutines*nIncrementos))

	fmt.Println("=== Variante INSEGURA:")
	res1 := incrementarInseguro(nGoroutines, nIncrementos)