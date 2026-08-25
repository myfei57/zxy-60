package sorter

import (
	"bagsort/internal/model"
)

func (s *Sorter) Divert(bag model.Bag) error {
	snap, err := s.store.Load()
	if err != nil {
		return err
	}
	bag.State = model.BagDiverted
	snap.Bags[bag.ID] = bag
	return s.store.Save(snap)
}

func (s *Sorter) ListRecords() ([]model.SortRecord, error) {
	snap, err := s.store.Load()
	if err != nil {
		return nil, err
	}
	return snap.SortRecords, nil
}
