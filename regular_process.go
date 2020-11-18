package main

func countPrimesRegular(inputStream <-chan int) <-chan int {
	outputStream := make(chan int, len(inputStream))
	defer close(outputStream)

	for i := range inputStream {
		if isPrime(i) {
			outputStream <- i
		}
	}

	return outputStream
}
