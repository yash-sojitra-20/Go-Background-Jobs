package workerpool

import "context"

type Job func(ctx context.Context)