package verifycase

import (
	"testing"

	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
)

func TestBsBagFlightMappingFresh(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	srt := sorter.NewSorter(st, book, chutes)

	if err := book.Register(model.Flight{ID: "F1", Code: "CA101", State: model.FlightOpen}); err != nil {
		t.Fatal(err)
	}
	if err := book.Register(model.Flight{ID: "F2", Code: "CA202", State: model.FlightOpen}); err != nil {
		t.Fatal(err)
	}
	if err := chutes.Assign("F1", "A"); err != nil {
		t.Fatal(err)
	}
	if err := chutes.Assign("F2", "B"); err != nil {
		t.Fatal(err)
	}
	bag := model.Bag{ID: "bag1", FlightID: "F1"}
	if err := book.AssignBag(bag.ID, "F1"); err != nil {
		t.Fatal(err)
	}
	if err := book.ScheduleChange(bag.ID, "F2"); err != nil {
		t.Fatal(err)
	}

	chuteID, err := srt.Route(bag)
	if err != nil {
		t.Fatal(err)
	}
	if chuteID != "B" {
		t.Fatalf("expected chute B from refreshed mapping, got %q", chuteID)
	}
}
