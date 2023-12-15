package examples

import (
	"time"
)

func proccessing() {
	// Setup the test array of numbers
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = i + 1
	}
	// Concurrent processing
	for i := 0; i < len(numbers); i++ {
		// sleep for 1 second to simulate processing
		time.Sleep(1 * time.Second)
	}

}

func main() {
	proccessing()
}
