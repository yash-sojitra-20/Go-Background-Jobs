package workerpool

import (
	"errors"
	"sync"
)

var ErrPoolClosed = errors.New("worker pool is closed")

type Pool struct {
	jobs chan *Job

	producerWG sync.WaitGroup
	workerWG   sync.WaitGroup

	shutdown chan struct{}
	once     sync.Once
}

func New(bufferSize int) *Pool {
	return &Pool{
		jobs:     make(chan *Job, bufferSize),
		shutdown: make(chan struct{}),
	}
}

func (p *Pool) Submit(job *Job) error {
	select {
	case <-p.shutdown:
		return ErrPoolClosed

	case p.jobs <- job:
		return nil
	}
}

func (p *Pool) Shutdown() {
	p.once.Do(func() {
		close(p.shutdown)
	})

	p.workerWG.Wait()
}

func (p *Pool) StartProducer(fn func()) {
	p.producerWG.Add(1)

	go func() {
		defer p.producerWG.Done()

		fn()
	}()
}
