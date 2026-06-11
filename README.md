# Go-Background-Jobs


Mini Background Job Engine:
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