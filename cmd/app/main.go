package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yash-sojitra-20/Go-Background-Jobs/internal/workerpool"
)

func main() {
	fmt.Println("Application Started")

	ctx, cancel := context.WithCancel(context.Background())

	pool := workerpool.New(10)

	pool.Start(ctx, 3)

	for i := 1; i <= 20; i++ {
		jobID := i

		pool.Submit(func(ctx context.Context) {
			fmt.Printf("processing job %d\n", jobID)

			time.Sleep(2 * time.Second)

			fmt.Printf("completed job %d\n", jobID)
		})
	}

	time.Sleep(5 * time.Second)

	fmt.Println("Initiating shutdown")

	cancel()

	pool.Shutdown()

	fmt.Println("Application shutdown complete")
}