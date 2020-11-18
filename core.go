package main

import "math"

func isPrime(number int) bool {
	if number == 1 {
		return false
	}
	if number == 2 {
		return true
	}
	if number%2 == 0 {
		return false
	}

	for d := 3; d <= int(math.Sqrt(float64(number))); d++ {
		if number%d == 0 {
			return false
		}
	}

	return true
}

func toStream(num int) <-chan int {
	stream := make(chan int, num)
	defer close(stream)

	for i := num; i > 0; i-- {
		stream <- i
	}

	return stream
}
