package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yash-sojitra-20/Go-Background-Jobs/internal/workerpool"
)

func main() {
	fmt.Println("Application Started")

	ctx, cancel := context.WithCancel(
		context.Background(),
	)

	defer cancel()

	pool := workerpool.New(10)

	pool.Start(ctx, 3)

	sigChan := make(chan os.Signal, 1)

	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		sig := <-sigChan

		fmt.Printf(
			"received signal: %v\n",
			sig,
		)

		fmt.Println(
			"starting graceful shutdown",
		)

		cancel()

		pool.Shutdown()

		fmt.Println(
			"application shutdown complete",
		)

		os.Exit(0)
	}()

	for i := 1; i <= 20; i++ {
		jobID := i

		job := &workerpool.Job{
			ID:     fmt.Sprintf("job-%d", jobID),
			Name:   "Demo Job",
			Status: workerpool.StatusPending,

			Handler: func(ctx context.Context) error {
				fmt.Printf(
					"processing %s\n",
					fmt.Sprintf("job-%d", jobID),
				)

				time.Sleep(2 * time.Second)

				fmt.Printf(
					"completed %s\n",
					fmt.Sprintf("job-%d", jobID),
				)

				return nil
			},
		}

		err := pool.Submit(job)

		if err != nil {
			fmt.Printf(
				"failed to submit %s: %v\n",
				job.ID,
				err,
			)
		}
	}

	select {}
}
