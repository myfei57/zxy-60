package sorter

import "bagsort/internal/model"

func (s *Sorter) Route(bag model.Bag) (string, error) {
	return s.chutes.Current(bag.FlightID)
}
