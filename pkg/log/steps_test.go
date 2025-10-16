package log

import (
	"fmt"
	"testing"
	"time"
)

func TestSteps(t *testing.T) {
	t.Skip("Skipping steps test")
	fmt.Println("Starting job test")
	renderer := NewRenderer()
	renderer.Start()

	// Simulate receiving job updates from a server
	for _, job := range mockJobUpdates() {
		renderer.Update(job.Steps)
		time.Sleep(2 * time.Second)
	}
	// renderer.PrintFinalSnapshot("job-123")
	renderer.Stop()

	renderer2 := NewRenderer()
	renderer2.Start()
	for _, job := range mockJobUpdates() {
		renderer2.Update(job.Steps)
		time.Sleep(2 * time.Second)
	}
	// renderer2.PrintFinalSnapshot("job-123")
	renderer2.Stop()

	fmt.Println("Job test completed")
}

// mockJobUpdates simulates job updates as if fetched from a server.
func mockJobUpdates() []Job {
	return []Job{
		{
			ID: "job-123",
			Steps: []Step{
				{Name: "Initialize", Status: "running"},
			},
		},
		{
			ID: "job-123",
			Steps: []Step{
				{Name: "Initialize", Status: "done"},
				{Name: "Download Data", Status: "running", Current: 30, Total: 100},
			},
		},
		{
			ID: "job-123",
			Steps: []Step{
				{Name: "Initialize", Status: "done"},
				{Name: "Download Data", Status: "running", Current: 80, Total: 100},
				{Name: "Train Model", Status: "running", Current: 10, Total: 200},
			},
		},
		{
			ID: "job-123",
			Steps: []Step{
				{Name: "Initialize", Status: "done"},
				{Name: "Download Data", Status: "done"},
				{Name: "Train Model", Status: "failed", Current: 190, Total: 200},
				{Name: "Deploy", Status: "pending"},
			},
		},
	}
}
