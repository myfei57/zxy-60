package flight

import (
	"errors"

	"bagsort/internal/model"
)

var ErrInvalidFlight = errors.New("invalid flight")

func (b *Book) Validate(f model.Flight) error {
	if err := model.ValidateFlight(f); err != nil {
		return ErrInvalidFlight
	}
	return nil
}

func (b *Book) Open(flightID string) error {
	return b.setState(flightID, model.FlightOpen)
}

func (b *Book) Board(flightID string) error {
	return b.setState(flightID, model.FlightBoarding)
}

func (b *Book) Cutoff(flightID string) error {
	return b.setState(flightID, model.FlightCutoff)
}

func (b *Book) setState(flightID string, state model.FlightState) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	f, ok := snap.Flights[flightID]
	if !ok {
		return ErrNotFound
	}
	f.State = state
	snap.Flights[flightID] = f
	return b.store.Save(snap)
}
