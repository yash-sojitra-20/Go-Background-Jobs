package runner

import "sync"

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

		task()
	}()
}

func (r *Runner) Wait() {
	r.wg.Wait()
}