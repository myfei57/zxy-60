package inject

import (
	"bagsort/internal/model"
)

func (inj *Injector) DrainInTransit(flightID string) error {
	bags := inj.inFlight[flightID]
	if len(bags) == 0 {
		inj.inFlight[flightID] = nil
		return nil
	}
	open := inj.book.IsOpen(flightID)
	snap, err := inj.store.Load()
	if err != nil {
		return err
	}
	for _, bag := range bags {
		if open {
			bag.State = model.BagLoaded
		} else {
			bag.State = model.BagDiverted
		}
		snap.Bags[bag.ID] = bag
	}
	inj.inFlight[flightID] = nil
	return inj.store.Save(snap)
}
