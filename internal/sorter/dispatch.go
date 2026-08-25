package sorter

import (
	"errors"

	"bagsort/internal/model"
	"bagsort/internal/tag"
)

var ErrReadNotCommitted = errors.New("barcode read is not committed")

func (s *Sorter) Dispatch(bag model.Bag, reading tag.Reading) error {
	chuteID, err := s.Route(bag)
	if err != nil {
		return err
	}
	return s.recordSort(bag, chuteID)
}

func (s *Sorter) Sort(bag model.Bag, reading tag.Reading) error {
	return s.Dispatch(bag, reading)
}
