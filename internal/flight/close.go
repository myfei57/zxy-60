package flight

import (
	"errors"

	"bagsort/internal/model"
)

// ErrFlightClosed is returned when an operation is rejected because the
// flight has already been closed. Once a flight is closed, check-in must
// not accept any further bags.
var ErrFlightClosed = errors.New("flight is closed")

type Drainer interface {
	DrainInTransit(flightID string) error
}

// Close drains the in-transit queue before marking the flight closed. The
// drain runs while the flight is still open so that in-transit bags are
// loaded rather than diverted. Only after the queue is drained is the
// flight marked closed, at which point check-in can no longer receive bags.
func (b *Book) Close(flightID string, drainer Drainer) error {
	if err := drainer.DrainInTransit(flightID); err != nil {
		return err
	}
	return b.markClosed(flightID)
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
