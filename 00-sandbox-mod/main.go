// B"H

/*
go mod init sandbox/concurrency
go run main.go
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	doneStream := make(chan interface{})

	go func() {
		time.Sleep(5 * time.Second)
		close(doneStream)
	}()

	workCounter := 0

loop:
	for {
		select {
		case <-doneStream:
			break loop
		default:
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("Achieved %v cycles of work before signalled to stop.\n", workCounter)
}

// RESULT: "Achieved 5 cycles of work before signalled to stop."
