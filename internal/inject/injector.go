package inject

import (
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
)

type Injector struct {
	store    *store.Store
	book     *flight.Book
	sorter   *sorter.Sorter
	inFlight map[string][]model.Bag
}

func NewInjector(st *store.Store, book *flight.Book, s *sorter.Sorter) *Injector {
	return &Injector{
		store:    st,
		book:     book,
		sorter:   s,
		inFlight: map[string][]model.Bag{},
	}
}

func (inj *Injector) Queue(flightID string) []model.Bag {
	return inj.inFlight[flightID]
}

func (inj *Injector) InTransit(flightID string) int {
	return len(inj.inFlight[flightID])
}
