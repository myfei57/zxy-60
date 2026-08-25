package sorter

import "bagsort/internal/model"

func (s *Sorter) Route(bag model.Bag) (string, error) {
	flightID, err := s.book.BagFlight(bag.ID)
	if err != nil {
		return "", err
	}
	return s.chutes.Current(flightID)
}
