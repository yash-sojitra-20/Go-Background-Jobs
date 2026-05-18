package main

import (
	"fmt"
	"time"

	"github.com/yash-sojitra-20/Go-Background-Jobs/internal/runner"
)

func main() {
	fmt.Println("Application Started")

	runner.Run(func() {
		fmt.Println("Background task started")

		time.Sleep(2 * time.Second)

		fmt.Println("Background task completed")
	})

	fmt.Println("Main function completed")

	time.Sleep(3 * time.Second)
}