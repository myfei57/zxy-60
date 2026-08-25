package sorter

import "bagsort/internal/model"

func (s *Sorter) RouteDetail(bag model.Bag) (model.RouteDetail, error) {
	flightID, err := s.book.BagFlight(bag.ID)
	if err != nil {
		return model.RouteDetail{}, err
	}
	chuteID, err := s.chutes.Current(flightID)
	if err != nil {
		return model.RouteDetail{}, err
	}
	return model.RouteDetail{BagID: bag.ID, FlightID: flightID, ChuteID: chuteID}, nil
}
