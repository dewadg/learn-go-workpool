package main

import (
	"sync"
)

func findPrimesWorkpool(inputStream <-chan int) <-chan int {
	outputStream := make(chan int, len(inputStream))

	wg := sync.WaitGroup{}
	wg.Add(len(inputStream))

	for i := 0; i < numOfWorkers; i++ {
		go worker(&wg, inputStream, outputStream)
	}

	wg.Wait()

	return outputStream
}

func worker(wg *sync.WaitGroup, inputStream <-chan int, outputStream chan int) {
	for i := range inputStream {
		if isPrime(i) {
			outputStream <- i
		}

		wg.Done()
	}
}
