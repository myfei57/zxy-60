package flight

import "bagsort/internal/model"

type Drainer interface {
	DrainInTransit(flightID string) error
}

func (b *Book) Close(flightID string, drainer Drainer) error {
	if err := b.markClosed(flightID); err != nil {
		return err
	}
	return drainer.DrainInTransit(flightID)
}

func (b *Book) markClosed(flightID string) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	f, ok := snap.Flights[flightID]
	if !ok {
		return ErrNotFound
	}
	f.State = model.FlightClosed
	snap.Flights[flightID] = f
	return b.store.Save(snap)
}
