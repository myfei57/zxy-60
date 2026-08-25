package tag

import (
	"bagsort/internal/model"
	"bagsort/internal/ns"
	"bagsort/internal/store"
)

type Reading struct {
	Barcode   string
	BagID     string
	Sequence  uint64
	Epoch     uint64
	Committed bool
}

type Reader struct {
	store *store.Store
	ns    *ns.Namespace
}

func NewReader(st *store.Store, namespace *ns.Namespace) *Reader {
	return &Reader{store: st, ns: namespace}
}

func (r *Reader) Read(barcode string) (Reading, error) {
	snap, err := r.store.Load()
	if err != nil {
		return Reading{}, err
	}
	return Reading{
		Barcode:   barcode,
		BagID:     r.ns.BagID(barcode),
		Sequence:  snap.Sequence.Next,
		Epoch:     snap.Sequence.Epoch,
		Committed: false,
	}, nil
}

// Commit persists the barcode read so that the sorter can verify the read was
// submitted before dispatching sort instructions. A read that has not been
// committed must never participate in scheduling.
func (r *Reader) Commit(reading Reading) error {
	snap, err := r.store.Load()
	if err != nil {
		return err
	}
	snap.Sequence = Allocate(snap.Sequence)
	bag := model.Bag{
		ID:       reading.BagID,
		Barcode:  reading.Barcode,
		Sequence: reading.Sequence,
		State:    model.BagCheckedIn,
	}
	snap.Bags[bag.ID] = bag
	if snap.CommittedReads == nil {
		snap.CommittedReads = map[string]bool{}
	}
	snap.CommittedReads[bag.ID] = true
	return r.store.Save(snap)
}

// IsCommitted reports whether the barcode read for bagID has been committed.
func (r *Reader) IsCommitted(bagID string) (bool, error) {
	snap, err := r.store.Load()
	if err != nil {
		return false, err
	}
	return snap.CommittedReads[bagID], nil
}
