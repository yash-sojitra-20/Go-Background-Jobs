package main

import (
	"fmt"
	"time"

	"github.com/yash-sojitra-20/Go-Background-Jobs/internal/runner"
)

func main() {
	fmt.Println("Application Started")

	r := runner.New()

	r.Run(func() {
		fmt.Println("Task 1 started")

		time.Sleep(1 * time.Second)

		panic("something went wrong in task 1")
	})

	r.Run(func() {
		fmt.Println("Task 2 started")

		time.Sleep(2 * time.Second)

		fmt.Println("Task 2 completed")
	})

	r.Wait()

	fmt.Println("Application shutdown gracefully")
}