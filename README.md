# Go-Background-Jobs


## Background Job Engine:

Features: 
```text
✓ Producer tracking
✓ Worker tracking
✓ Queue ownership
✓ Panic recovery
✓ Job lifecycle
✓ Graceful shutdown
✓ Deterministic draining
```

The shutdown sequence:
```
CTRL+C
↓
close(shutdown)
↓
producerWG.Wait()
↓
close(jobs)
↓
workers drain queue
↓
workerWG.Wait()
↓
exit
```

---

# Run the engine: go run cmd/api/main.go

## Scenario 1: Shutdown after all jobs completed

Log:

```text
completed job-20

^C
received signal: interrupt
starting graceful shutdown

worker 2 shutting down
worker 1 shutting down
worker 3 shutting down

application shutdown complete
```

### Verdict

```text
PASS ✅
```

Reason:

```text
Queue already empty
↓
Workers immediately exit
↓
Shutdown completes
```

---

## Scenario 2: Shutdown during producer execution

Log:

```text
^C
received signal: interrupt
starting graceful shutdown

failed to submit job-14: worker pool is closed
```

This means:

```text
Shutdown initiated
↓
New submissions rejected
↓
Producer stopped
```

### Verdict

```text
PASS ✅
```

---

## Scenario 3: Drain accepted work

After shutdown:

```text
completed job-2
processing job-4

completed job-1
processing job-5

completed job-3
processing job-6
...
processing job-13
```

Notice:

```text
Jobs already accepted
were NOT discarded.
```

### Verdict

```text
PASS ✅
```

The warm shutdown policy.

---

## Scenario 4: Worker Exit

This was the hardest problem.

Example:

```text
completed job-10
worker 2 shutting down

completed job-11
worker 1 shutting down

completed job-13
worker 3 shutting down
```

At first glance, this may look random. But it's actually deterministic.

Why?

Because:

```text
Worker exits
ONLY AFTER
it finishes draining its share of accepted jobs.
```

The exact order depends on:

```text
Which worker happened to receive which jobs.
```

NOT on:

```text
Random select between ctx.Done() and jobs.
```

### Verdict

```text
PASS ✅
```

---
