package verifycase

import (
	"testing"

	"bagsort/internal/model"
	"bagsort/internal/recheck"
)

func TestBsRecheckQueueOrder(t *testing.T) {
	q := recheck.NewQueue()
	q.Enqueue(model.Bag{ID: "bag1"})
	q.Enqueue(model.Bag{ID: "bag2"})

	first, ok := q.Dequeue()
	if !ok {
		t.Fatal("queue is empty")
	}
	if first.ID != "bag1" {
		t.Fatalf("expected FIFO bag1, got %s", first.ID)
	}
}
