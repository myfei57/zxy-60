package inject

import (
	"bagsort/internal/model"
	"bagsort/internal/tag"
)

func (inj *Injector) Push(bag model.Bag, chuteID string) error {
	if err := inj.sorter.Dispatch(bag, tag.Reading{Committed: true}); err != nil {
		return err
	}
	return inj.sorter.AssignChute(bag.FlightID, chuteID)
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
