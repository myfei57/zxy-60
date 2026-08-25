package sorter

import (
	"errors"

	"bagsort/internal/model"
	"bagsort/internal/tag"
)

// ErrReadNotCommitted is returned when a sort instruction is requested for a
// barcode read that has not been submitted (committed). Such a read carries no
// authoritative identity, so dispatching would route the bag based on stale data.
var ErrReadNotCommitted = errors.New("barcode read is not committed")

// Dispatch issues a sort instruction for bag, but only after confirming that
// the barcode read has been committed. A read that has not been submitted must
// never participate in scheduling, otherwise a failed read inherits the
// previous bag's route and is diverted to the wrong chute.
func (s *Sorter) Dispatch(bag model.Bag, reading tag.Reading) error {
	// The reading argument is the caller's request to dispatch; the source of
	// truth for whether the read was actually committed is the persisted commit
	// set, not the in-memory flag the caller hands us. This prevents a stale or
	// failed read from being dispatched as if it were the previous bag.
	committed, err := s.readCommitted(bag)
	if err != nil {
		return err
	}
	if !committed {
		return ErrReadNotCommitted
	}
	chuteID, err := s.Route(bag)
	if err != nil {
		return err
	}
	return s.recordSort(bag, chuteID)
}

// Sort is an alias for Dispatch retained for backwards compatibility.
func (s *Sorter) Sort(bag model.Bag, reading tag.Reading) error {
	return s.Dispatch(bag, reading)
}
