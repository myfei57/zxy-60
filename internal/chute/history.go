package chute

import (
	"strings"

	"bagsort/internal/model"
)

func (a *Assigner) List() ([]model.ChuteAssignment, error) {
	snap, err := a.store.Load()
	if err != nil {
		return nil, err
	}
	out := make([]model.ChuteAssignment, 0, len(snap.ChuteAssignments))
	for flightID, chuteID := range snap.ChuteAssignments {
		out = append(out, model.ChuteAssignment{FlightID: flightID, ChuteID: chuteID})
	}
	return out, nil
}

func (a *Assigner) Validate(chuteID string) bool {
	return strings.TrimSpace(chuteID) != ""
}
