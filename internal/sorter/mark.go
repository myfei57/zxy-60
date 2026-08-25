package sorter

import (
	"time"

	"bagsort/internal/model"
)

func (s *Sorter) recordSort(bag model.Bag, chuteID string) error {
	snap, err := s.store.Load()
	if err != nil {
		return err
	}
	snap.SortRecords = append(snap.SortRecords, model.SortRecord{
		BagID:    bag.ID,
		FlightID: bag.FlightID,
		ChuteID:  chuteID,
		Sequence: bag.Sequence,
		At:       time.Now().UTC().Format(time.RFC3339),
	})
	return s.store.Save(snap)
}

func (s *Sorter) WasExecuted(bagID string) (bool, error) {
	snap, err := s.store.Load()
	if err != nil {
		return false, err
	}
	for _, record := range snap.SortRecords {
		if record.BagID == bagID {
			return true, nil
		}
	}
	return false, nil
}

func (s *Sorter) SortCount(bagID string) (int, error) {
	snap, err := s.store.Load()
	if err != nil {
		return 0, err
	}
	count := 0
	for _, record := range snap.SortRecords {
		if record.BagID == bagID {
			count++
		}
	}
	return count, nil
}

func (s *Sorter) LastChute(bagID string) (string, error) {
	snap, err := s.store.Load()
	if err != nil {
		return "", err
	}
	var last string
	for _, record := range snap.SortRecords {
		if record.BagID == bagID {
			last = record.ChuteID
		}
	}
	return last, nil
}
