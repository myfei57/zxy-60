package store

import "bagsort/internal/model"

func (s *Store) Stats() (model.SnapshotStats, error) {
	snap, err := s.Load()
	if err != nil {
		return model.SnapshotStats{}, err
	}
	return model.SnapshotStats{
		Flights:          len(snap.Flights),
		Bags:             len(snap.Bags),
		ChuteAssignments: len(snap.ChuteAssignments),
		SortRecords:      len(snap.SortRecords),
		TransferCommands: len(snap.TransferCommands),
	}, nil
}

func (s *Store) Clear() error {
	return s.Save(model.NewSnapshot())
}
