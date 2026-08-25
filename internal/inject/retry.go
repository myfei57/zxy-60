package inject

import (
	"bagsort/internal/model"
	"bagsort/internal/tag"
)

func (inj *Injector) Retry(bag model.Bag) error {
	done, err := inj.sorter.WasExecuted(bag.ID)
	if err != nil {
		return err
	}
	if done {
		return nil
	}
	return inj.sorter.Sort(bag, tag.Reading{Committed: true})
}
