package store

import (
	"sort"

	"bagsort/internal/model"
)

func (s *Store) GetBag(bagID string) (model.Bag, bool, error) {
	snap, err := s.Load()
	if err != nil {
		return model.Bag{}, false, err
	}
	bag, ok := snap.Bags[bagID]
	return bag, ok, nil
}

func (s *Store) PutBag(bag model.Bag) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	snap.Bags[bag.ID] = bag
	return s.Save(snap)
}

func (s *Store) ListBags() ([]model.Bag, error) {
	snap, err := s.Load()
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(snap.Bags))
	for id := range snap.Bags {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	out := make([]model.Bag, 0, len(ids))
	for _, id := range ids {
		out = append(out, snap.Bags[id])
	}
	return out, nil
}

func (s *Store) GetFlight(flightID string) (model.Flight, bool, error) {
	snap, err := s.Load()
	if err != nil {
		return model.Flight{}, false, err
	}
	f, ok := snap.Flights[flightID]
	return f, ok, nil
}

func (s *Store) PutFlight(f model.Flight) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	snap.Flights[f.ID] = f
	return s.Save(snap)
}

func (s *Store) ListFlights() ([]model.Flight, error) {
	snap, err := s.Load()
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

func (s *Store) PutChuteAssignment(flightID string, chuteID string) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	snap.ChuteAssignments[flightID] = chuteID
	return s.Save(snap)
}

func (s *Store) GetChuteAssignment(flightID string) (string, bool, error) {
	snap, err := s.Load()
	if err != nil {
		return "", false, err
	}
	chuteID, ok := snap.ChuteAssignments[flightID]
	return chuteID, ok, nil
}

func (s *Store) ListChuteAssignments() ([]model.ChuteAssignment, error) {
	snap, err := s.Load()
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(snap.ChuteAssignments))
	for id := range snap.ChuteAssignments {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	out := make([]model.ChuteAssignment, 0, len(ids))
	for _, id := range ids {
		out = append(out, model.ChuteAssignment{FlightID: id, ChuteID: snap.ChuteAssignments[id]})
	}
	return out, nil
}

func (s *Store) AddSortRecord(record model.SortRecord) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	snap.SortRecords = append(snap.SortRecords, record)
	return s.Save(snap)
}

func (s *Store) ListSortRecords() ([]model.SortRecord, error) {
	snap, err := s.Load()
	if err != nil {
		return nil, err
	}
	return snap.SortRecords, nil
}

func (s *Store) PutSequence(seq model.Sequence) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	snap.Sequence = seq
	return s.Save(snap)
}

func (s *Store) Sequence() (model.Sequence, error) {
	snap, err := s.Load()
	if err != nil {
		return model.Sequence{}, err
	}
	return snap.Sequence, nil
}
