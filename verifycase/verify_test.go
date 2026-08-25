package verifycase

import (
	"testing"

	"bagsort/internal/belt"
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/store"
)

func TestBsBatchTransferOrder(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	b := belt.NewBelt(st, book)
	b.SetCarousel("A", 1)
	b.Place(model.Bag{ID: "bag1", FlightID: "F1"})

	if err := b.Switch("B", 2); err != nil {
		t.Fatal(err)
	}
	loads := b.Loads()
	if len(loads) != 1 {
		t.Fatalf("expected 1 load record, got %d", len(loads))
	}
	if loads[0].Carousel != "A" || loads[0].Batch != 1 {
		t.Fatalf("expected load on carousel A batch 1, got %s/%d", loads[0].Carousel, loads[0].Batch)
	}
}
