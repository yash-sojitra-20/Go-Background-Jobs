package runner

import (
	"log"
	"sync"
)

type Runner struct {
	wg sync.WaitGroup
}

func New() *Runner {
	return &Runner{}
}

func (r *Runner) Run(task func()) {
	r.wg.Add(1)

	go func() {
		defer r.wg.Done()

		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic recovered: %v", err)
			}
		}()

		task()
	}()
}

func (r *Runner) Wait() {
	r.wg.Wait()
}