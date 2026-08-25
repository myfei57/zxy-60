package inject

import (
	"bagsort/internal/model"
	"bagsort/internal/tag"
)

func (inj *Injector) Retry(bag model.Bag) error {
	return inj.sorter.Sort(bag, tag.Reading{Committed: true})
}
