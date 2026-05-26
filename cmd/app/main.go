package main

import (
	"context"
	"fmt"
	"time"

	"github.com/yash-sojitra-20/Go-Background-Jobs/internal/runner"
)

func main() {
	fmt.Println("Application Started")

	r := runner.New()

	ctx, cancel := context.WithCancel(context.Background())

	r.Run(ctx, func(ctx context.Context) {
		fmt.Println("Task started")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Task received cancellation signal")
				return

			default:
				fmt.Println("Task working...")

				time.Sleep(1 * time.Second)
			}
		}
	})

	time.Sleep(5 * time.Second)

	fmt.Println("Cancelling context")

	cancel()

	r.Wait()

	fmt.Println("Application shutdown gracefully")
}