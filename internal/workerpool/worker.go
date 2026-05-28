package workerpool

import (
	"context"
	"log"
)

func (p *Pool) Start(ctx context.Context, workerCount int) {
	for i := 1; i <= workerCount; i++ {
		p.wg.Add(1)

		go p.worker(ctx, i)
	}
}

func (p *Pool) worker(ctx context.Context, id int) {
	defer p.wg.Done()

	log.Printf("worker %d started", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d shutting down", id)
			return

		case job, ok := <-p.jobs:
			if !ok {
				log.Printf("worker %d job channel closed", id)
				return
			}

			p.executeJob(ctx, job, id)
		}
	}
}

func (p *Pool) executeJob(
	ctx context.Context,
	job Job,
	workerID int,
) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf(
				"worker %d recovered panic: %v",
				workerID,
				err,
			)
		}
	}()

	job(ctx)
}