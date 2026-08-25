package verifycase

import (
	"testing"

	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/inject"
	"bagsort/internal/model"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
)

func TestBsInjectAfterChutePersist(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	srt := sorter.NewSorter(st, book, chutes)
	inj := inject.NewInjector(st, book, srt)

	if err := book.Register(model.Flight{ID: "F1", Code: "CA101", State: model.FlightOpen}); err != nil {
		t.Fatal(err)
	}
	if err := srt.AssignChute("F1", "A"); err != nil {
		t.Fatal(err)
	}
	bag := model.Bag{ID: "bag1", Barcode: "BC001", FlightID: "F1"}
	if err := book.AssignBag(bag.ID, "F1"); err != nil {
		t.Fatal(err)
	}

	if err := inj.Push(bag, "B"); err != nil {
		t.Fatal(err)
	}
	chuteID, err := srt.LastChute(bag.ID)
	if err != nil {
		t.Fatal(err)
	}
	if chuteID != "B" {
		t.Fatalf("expected bag sorted to chute B, got %q", chuteID)
	}
}
