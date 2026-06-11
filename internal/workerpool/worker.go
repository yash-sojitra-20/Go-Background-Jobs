package workerpool

import (
	"context"
	"log"
)

func (p *Pool) Start(
	ctx context.Context,
	workerCount int,
) {
	for i := 1; i <= workerCount; i++ {
		p.workerWG.Add(1)

		go p.worker(ctx, i)
	}
}

func (p *Pool) worker(
	ctx context.Context,
	id int,
) {
	defer p.workerWG.Done()

	log.Printf("worker %d started", id)

	for {
		select {
		case <-ctx.Done():
			log.Printf(
				"worker %d shutting down",
				id,
			)

			return

		case job := <-p.jobs:
			job.Status = StatusRunning

			p.executeJob(
				ctx,
				job,
				id,
			)
		}
	}
}

func (p *Pool) executeJob(
	ctx context.Context,
	job *Job,
	workerID int,
) {
	defer func() {
		if err := recover(); err != nil {
			job.Status = StatusFailed

			log.Printf(
				"worker %d recovered panic in job %s: %v",
				workerID,
				job.ID,
				err,
			)
		}
	}()

	if err := job.Handler(ctx); err != nil {
		job.Status = StatusFailed

		log.Printf(
			"worker %d job %s failed: %v",
			workerID,
			job.ID,
			err,
		)

		return
	}

	job.Status = StatusCompleted
}
