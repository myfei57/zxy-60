package store

import (
	"encoding/json"
	"os"
	"path/filepath"

	"bagsort/internal/model"
)

type Store struct {
	dir  string
	name string
}

func New(dir string, name string) *Store {
	return &Store{dir: dir, name: name}
}

func (s *Store) Path() string {
	return filepath.Join(s.dir, s.name)
}

func (s *Store) Load() (*model.Snapshot, error) {
	data, err := os.ReadFile(s.Path())
	if err != nil {
		if os.IsNotExist(err) {
			return model.NewSnapshot(), nil
		}
		return nil, err
	}
	var snap model.Snapshot
	if err := json.Unmarshal(data, &snap); err != nil {
		return nil, err
	}
	if snap.Flights == nil {
		snap.Flights = map[string]model.Flight{}
	}
	if snap.Bags == nil {
		snap.Bags = map[string]model.Bag{}
	}
	if snap.FlightMappings == nil {
		snap.FlightMappings = map[string]string{}
	}
	if snap.ChuteAssignments == nil {
		snap.ChuteAssignments = map[string]string{}
	}
	if snap.TransferCommands == nil {
		snap.TransferCommands = []model.TransferCommand{}
	}
	if snap.SortRecords == nil {
		snap.SortRecords = []model.SortRecord{}
	}
	return &snap, nil
}

func (s *Store) Save(snap *model.Snapshot) error {
	if err := os.MkdirAll(s.dir, 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(snap, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.Path() + ".tmp"
	if err := os.WriteFile(tmp, data, 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, s.Path())
}
