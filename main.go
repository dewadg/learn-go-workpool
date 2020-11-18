package main

import (
	"fmt"
	"time"
)

const numOfWorkers = 2
const max = 10000000

func main() {
	input := toStream(max)
	start := time.Now()
	fmt.Println("Processing without workpool")
	fmt.Printf("Count\t\t: %d\n", len(countPrimesRegular(input)))
	fmt.Printf("Duration\t: %s\n", time.Since(start))

	fmt.Println()
	input = toStream(max)
	start = time.Now()
	fmt.Println("Processing with workpool")
	fmt.Printf("Count\t\t: %d\n", len(countPrimesWorkpool(input)))
	fmt.Printf("Duration\t: %s\n", time.Since(start))
}
