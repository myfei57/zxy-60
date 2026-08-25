package sorter

import (
	"errors"
	"testing"

	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/ns"
	"bagsort/internal/store"
	"bagsort/internal/tag"
)

// newTestSorter builds a sorter backed by a temp file store so that the
// persisted commit set drives the dispatch gate, exactly as in production.
func newTestSorter(t *testing.T) (*Sorter, *tag.Reader, *chute.Assigner, *flight.Book, *store.Store) {
	t.Helper()
	dir := t.TempDir()
	st := store.New(dir, "bagsort.json")
	namespace := ns.New("T1")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	reader := tag.NewReader(st, namespace)
	return NewSorter(st, book, chutes), reader, chutes, book, st
}

func TestDispatchRejectsUncommittedRead(t *testing.T) {
	s, reader, chutes, book, _ := newTestSorter(t)

	// Bag A checks in: read is committed, a chute is assigned, then dispatched.
	bagA := model.Bag{ID: "T1-A", Barcode: "A", Sequence: 1, FlightID: "F1"}
	if err := reader.Commit(tag.Reading{Barcode: "A", BagID: "T1-A", Sequence: 1}); err != nil {
		t.Fatalf("commit A: %v", err)
	}
	if err := book.AssignBag("T1-A", "F1"); err != nil {
		t.Fatalf("assign bag A to flight: %v", err)
	}
	if err := chutes.Assign("F1", "C1"); err != nil {
		t.Fatalf("assign chute: %v", err)
	}
	if err := s.Dispatch(bagA, tag.Reading{BagID: "T1-A"}); err != nil {
		t.Fatalf("dispatch committed read A: %v", err)
	}

	// Bag B's read failed and was never committed. Dispatch must refuse it so
	// it cannot inherit bag A's route and be sent to the wrong chute.
	bagB := model.Bag{ID: "T1-B", Barcode: "B", Sequence: 2, FlightID: "F1"}
	if err := book.AssignBag("T1-B", "F1"); err != nil {
		t.Fatalf("assign bag B to flight: %v", err)
	}
	err := s.Dispatch(bagB, tag.Reading{BagID: "T1-B"})
	if !errors.Is(err, ErrReadNotCommitted) {
		t.Fatalf("expected ErrReadNotCommitted for uncommitted read, got %v", err)
	}

	// No sort record may exist for the uncommitted bag.
	count, err := s.SortCount("T1-B")
	if err != nil {
		t.Fatalf("sort count: %v", err)
	}
	if count != 0 {
		t.Fatalf("uncommitted read produced %d sort record(s), want 0", count)
	}
}

func TestDispatchAllowsCommittedRead(t *testing.T) {
	s, reader, chutes, book, _ := newTestSorter(t)

	bag := model.Bag{ID: "T1-X", Barcode: "X", Sequence: 7, FlightID: "F1"}
	if err := reader.Commit(tag.Reading{Barcode: "X", BagID: "T1-X", Sequence: 7}); err != nil {
		t.Fatalf("commit: %v", err)
	}
	if err := book.AssignBag("T1-X", "F1"); err != nil {
		t.Fatalf("assign bag to flight: %v", err)
	}
	if err := chutes.Assign("F1", "C9"); err != nil {
		t.Fatalf("assign chute: %v", err)
	}
	if err := s.Dispatch(bag, tag.Reading{BagID: "T1-X"}); err != nil {
		t.Fatalf("dispatch committed read: %v", err)
	}
	last, err := s.LastChute("T1-X")
	if err != nil {
		t.Fatalf("last chute: %v", err)
	}
	if last != "C9" {
		t.Fatalf("dispatched to %q, want C9", last)
	}
}

func TestDispatchRejectsSyntheticCommittedFlag(t *testing.T) {
	s, _, chutes, _, _ := newTestSorter(t)

	// A caller must not be able to bypass the gate by fabricating a committed
	// reading for a bag whose read was never actually submitted.
	bag := model.Bag{ID: "T1-GHOST", Barcode: "G", FlightID: "F1"}
	if err := chutes.Assign("F1", "C1"); err != nil {
		t.Fatalf("assign chute: %v", err)
	}
	err := s.Dispatch(bag, tag.Reading{BagID: "T1-GHOST", Committed: true})
	if !errors.Is(err, ErrReadNotCommitted) {
		t.Fatalf("expected gate to ignore the in-memory flag, got %v", err)
	}
}
