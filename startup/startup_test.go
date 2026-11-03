package startup

import "testing"

func TestLongRunningWorkerPoolSize(t *testing.T) {
	const longRunningTasks = 7

	if longRunningWorkerPoolSize < longRunningTasks {
		t.Fatalf("long-running worker pool size = %d, want at least %d", longRunningWorkerPoolSize, longRunningTasks)
	}
}
