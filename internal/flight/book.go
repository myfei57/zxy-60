package flight

import (
	"errors"
	"sort"

	"bagsort/internal/model"
	"bagsort/internal/store"
)

var ErrNotFound = errors.New("flight not found")

type Book struct {
	store *store.Store
}

func NewBook(st *store.Store) *Book {
	return &Book{store: st}
}

func (b *Book) Register(f model.Flight) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	snap.Flights[f.ID] = f
	return b.store.Save(snap)
}

func (b *Book) Get(flightID string) (model.Flight, error) {
	snap, err := b.store.Load()
	if err != nil {
		return model.Flight{}, err
	}
	f, ok := snap.Flights[flightID]
	if !ok {
		return model.Flight{}, ErrNotFound
	}
	return f, nil
}

func (b *Book) List() ([]model.Flight, error) {
	snap, err := b.store.Load()
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(snap.Flights))
	for id := range snap.Flights {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	out := make([]model.Flight, 0, len(ids))
	for _, id := range ids {
		out = append(out, snap.Flights[id])
	}
	return out, nil
}

func (b *Book) IsOpen(flightID string) bool {
	f, err := b.Get(flightID)
	if err != nil {
		return false
	}
	return f.State == model.FlightOpen || f.State == model.FlightBoarding
}

func (b *Book) State(flightID string) (model.FlightState, error) {
	f, err := b.Get(flightID)
	if err != nil {
		return "", err
	}
	return f.State, nil
}
