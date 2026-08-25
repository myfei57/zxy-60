package verifycase

import (
	"testing"

	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/inject"
	"bagsort/internal/model"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
	"bagsort/internal/tag"
)

func TestBsSorterNoDuplicate(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	srt := sorter.NewSorter(st, book, chutes)
	inj := inject.NewInjector(st, book, srt)

	if err := book.Register(model.Flight{ID: "F1", Code: "CA101", State: model.FlightOpen}); err != nil {
		t.Fatal(err)
	}
	if err := chutes.Assign("F1", "A"); err != nil {
		t.Fatal(err)
	}
	bag := model.Bag{ID: "bag1", FlightID: "F1"}
	if err := book.AssignBag(bag.ID, "F1"); err != nil {
		t.Fatal(err)
	}

	if err := srt.Sort(bag, tag.Reading{Committed: true}); err != nil {
		t.Fatal(err)
	}
	if err := inj.Retry(bag); err != nil {
		t.Fatal(err)
	}
	count, err := srt.SortCount(bag.ID)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Fatalf("expected 1 sort record, got %d", count)
	}
}
