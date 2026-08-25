package inject

import (
	"bagsort/internal/model"
	"bagsort/internal/tag"
)

// Push assigns chuteID to the bag's flight and dispatches a sort instruction.
// Dispatch verifies that the bag's barcode read has been committed before
// issuing the instruction; an uncommitted read is rejected here so it cannot
// inherit the previous bag's route.
func (inj *Injector) Push(bag model.Bag, chuteID string) error {
	if err := inj.sorter.AssignChute(bag.FlightID, chuteID); err != nil {
		return err
	}
	return inj.sorter.Dispatch(bag, tag.Reading{BagID: bag.ID})
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
