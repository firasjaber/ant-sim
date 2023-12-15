package examples

import (
	"log"
	"sync"
	"time"
)

func concurrentProcessing() {
	// DEBUG: Measure execution time
	start := time.Now()
	defer func() {
		elapsed := time.Since(start)
		log.Printf("all processing took %s", elapsed)
	}()
	// Setup the test array of numbers
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}
	// Concurrent processing
	numRoutines := 4

	var wg sync.WaitGroup

	totalLength := len(numbers)
	chunkSize := (len(numbers) + numRoutines - 1) / numRoutines

	for i := 0; i < totalLength; i += chunkSize {
		// chunking
		end := i + chunkSize
		if end > totalLength {
			end = totalLength
		}

		// process chunk
		wg.Add(1)
		go processChunk(numbers, i, end, &wg)
	}

	wg.Wait()
}

func processChunk(numbers []int, start, end int, wg *sync.WaitGroup) {
	log.Println("process chunk started")

	for i := start; i < end; i++ {
		// sleep for 1 second to simulate processing
		time.Sleep(1 * time.Second)
	}
	wg.Done()
	log.Println("process chunk ended")
}

func main() {
	concurrentProcessing()
}
