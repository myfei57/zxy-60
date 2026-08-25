package inject

import (
	"bagsort/internal/model"
	"bagsort/internal/tag"
)

// Retry re-dispatches a sort instruction for a bag whose read has already been
// committed during check-in. Dispatch re-verifies the committed state, so an
// uncommitted read is rejected rather than inheriting a stale route.
func (inj *Injector) Retry(bag model.Bag) error {
	done, err := inj.sorter.WasExecuted(bag.ID)
	if err != nil {
		return err
	}
	if done {
		return nil
	}
	return inj.sorter.Sort(bag, tag.Reading{BagID: bag.ID})
}
