package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {

		// This case runs when cancel() is called
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return

		// This runs while the context is active
		default:
			fmt.Println("Working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {

	// Create a context with a cancel function
	ctx, cancel := context.WithCancel(context.Background())

	// Start the worker goroutine
	go worker(ctx)

	// Let the worker run for 5 seconds
	time.Sleep(5 * time.Second)

	fmt.Println("Main: Cancelling worker...")

	// Stop the worker
	cancel()

	// Give the worker time to print and exit
	time.Sleep(2 * time.Second)

	fmt.Println("Main: Program finished")
}
