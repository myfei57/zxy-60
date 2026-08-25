package inject

import (
	"bagsort/internal/model"
	"bagsort/internal/tag"
)

func (inj *Injector) Push(bag model.Bag, chuteID string) error {
	// Persist the chute assignment before dispatching the bag into the sort
	// queue. If the assignment is saved only after dispatch, a restart in
	// between recovers the stale (old) chute from the snapshot and the bag
	// is routed to the wrong flight. Routing must follow the latest
	// allocation after recovery.
	if err := inj.sorter.AssignChute(bag.FlightID, chuteID); err != nil {
		return err
	}
	return inj.sorter.Dispatch(bag, tag.Reading{Committed: true})
}

func (inj *Injector) Accept(bag model.Bag) error {
	inj.inFlight[bag.FlightID] = append(inj.inFlight[bag.FlightID], bag)
	snap, err := inj.store.Load()
	if err != nil {
		return err
	}
	bag.State = model.BagInjected
	snap.Bags[bag.ID] = bag
	return inj.store.Save(snap)
}
