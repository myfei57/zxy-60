package sorter

import (
	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/store"
)

type Sorter struct {
	store     *store.Store
	book      *flight.Book
	chutes    *chute.Assigner
	lastEpoch uint64
	lastSeq   uint64
}

func NewSorter(st *store.Store, book *flight.Book, chutes *chute.Assigner) *Sorter {
	return &Sorter{
		store:  st,
		book:   book,
		chutes: chutes,
	}
}

// readCommitted reports whether the barcode read for bag has been committed.
// Dispatch uses this as the gate: a read that has not been submitted must not
// participate in scheduling, otherwise a failed read inherits the previous
// bag's route and is diverted to the wrong chute.
func (s *Sorter) readCommitted(bag model.Bag) (bool, error) {
	snap, err := s.store.Load()
	if err != nil {
		return false, err
	}
	return snap.CommittedReads[bag.ID], nil
}
