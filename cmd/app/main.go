package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yash-sojitra-20/Go-Background-Jobs/internal/runner"
)

func main() {
	fmt.Println("Application Started")

	r := runner.New()

	ctx, cancel := context.WithCancel(context.Background())

	r.Run(ctx, func(ctx context.Context) {
		fmt.Println("Worker started")

		for {
			select {
			case <-ctx.Done():
				fmt.Println("Worker shutting down gracefully")
				return

			default:
				fmt.Println("Worker processing job")

				time.Sleep(2 * time.Second)
			}
		}
	})

	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan

	fmt.Printf("Received signal: %v\n", sig)

	fmt.Println("Initiating graceful shutdown")

	cancel()

	r.Wait()

	fmt.Println("Application shutdown complete")
}