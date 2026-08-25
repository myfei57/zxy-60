package chute

import (
	"bagsort/internal/store"
)

type Assigner struct {
	store *store.Store
}

func NewAssigner(st *store.Store) *Assigner {
	return &Assigner{store: st}
}

func (a *Assigner) Assign(flightID string, chuteID string) error {
	snap, err := a.store.Load()
	if err != nil {
		return err
	}
	snap.ChuteAssignments[flightID] = chuteID
	return a.store.Save(snap)
}

func (a *Assigner) Current(flightID string) (string, error) {
	snap, err := a.store.Load()
	if err != nil {
		return "", err
	}
	return snap.ChuteAssignments[flightID], nil
}
