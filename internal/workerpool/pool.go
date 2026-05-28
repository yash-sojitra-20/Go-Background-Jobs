package workerpool

import (
	"sync"
)

type Pool struct {
	jobs chan Job
	wg   sync.WaitGroup
}

func New(bufferSize int) *Pool {
	return &Pool{
		jobs: make(chan Job, bufferSize),
	}
}

func (p *Pool) Submit(job Job) {
	p.jobs <- job
}

func (p *Pool) Shutdown() {
	close(p.jobs)

	p.wg.Wait()
}