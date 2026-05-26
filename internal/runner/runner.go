package runner

import (
	"context"
	"log"
	"sync"
)

type Task func(ctx context.Context)

type Runner struct {
	wg sync.WaitGroup
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(ctx context.Context, task Task) {
	r.wg.Add(1)

	go func() {
		defer r.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v", err)
			}
		}()

		task(ctx)
	}()
}

func (r *Runner) Wait() {
	r.wg.Wait()
}