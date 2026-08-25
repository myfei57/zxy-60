package verifycase

import (
	"testing"

	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
	"bagsort/internal/tag"
)

func TestBsTagDispatchOrder(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	srt := sorter.NewSorter(st, book, chutes)

	if err := book.Register(model.Flight{ID: "F1", Code: "CA101", State: model.FlightOpen}); err != nil {
		t.Fatal(err)
	}
	if err := srt.AssignChute("F1", "A"); err != nil {
		t.Fatal(err)
	}
	bag := model.Bag{ID: "bag1", FlightID: "F1"}
	if err := book.AssignBag(bag.ID, "F1"); err != nil {
		t.Fatal(err)
	}

	if err := srt.Dispatch(bag, tag.Reading{Committed: false}); err == nil {
		t.Fatal("expected dispatch to fail when barcode read is not committed")
	}
}
