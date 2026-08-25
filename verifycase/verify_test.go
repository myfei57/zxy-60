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

func TestBsFlightCloseOrder(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	srt := sorter.NewSorter(st, book, chutes)
	inj := inject.NewInjector(st, book, srt)

	if err := book.Register(model.Flight{ID: "F1", Code: "CA101", State: model.FlightOpen}); err != nil {
		t.Fatal(err)
	}
	bag1 := model.Bag{ID: "bag1", Barcode: "BC001", FlightID: "F1"}
	bag2 := model.Bag{ID: "bag2", Barcode: "BC002", FlightID: "F1"}
	if err := inj.Accept(bag1); err != nil {
		t.Fatal(err)
	}
	if err := inj.Accept(bag2); err != nil {
		t.Fatal(err)
	}

	if err := book.Close("F1", inj); err != nil {
		t.Fatal(err)
	}
	got, _, err := st.GetBag("bag1")
	if err != nil {
		t.Fatal(err)
	}
	if got.State != model.BagLoaded {
		t.Fatalf("expected bag1 loaded, got %s", got.State)
	}
}
