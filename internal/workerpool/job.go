package workerpool

import (
	"context"
)

type Status string

const (
	StatusPending   Status = "pending"
	StatusRunning   Status = "running"
	StatusCompleted Status = "completed"
	StatusFailed    Status = "failed"
)

type Handler func(ctx context.Context) error

type Job struct {
	ID      string
	Name    string
	Status  Status
	Handler Handler
}
