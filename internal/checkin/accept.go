package checkin

import (
	"bagsort/internal/flight"
	"bagsort/internal/model"
)

func (d *Desk) CheckIn(barcode string, flightID string) (model.Bag, error) {
	// Reject bags for closed flights before any side effects: a closed flight
	// has finished loading, so late bags cannot make the flight.
	if d.book.IsClosed(flightID) {
		return model.Bag{}, flight.ErrFlightClosed
	}
	reading, err := d.reader.Read(barcode)
	if err != nil {
		return model.Bag{}, err
	}
	if err := d.reader.Commit(reading); err != nil {
		return model.Bag{}, err
	}
	bag := model.Bag{
		ID:       reading.BagID,
		Barcode:  barcode,
		Sequence: reading.Sequence,
		FlightID: flightID,
		State:    model.BagCheckedIn,
	}
	if err := model.ValidateBag(bag); err != nil {
		return model.Bag{}, err
	}
	if err := d.book.AssignBag(bag.ID, flightID); err != nil {
		return model.Bag{}, err
	}
	if err := d.injector.Accept(bag); err != nil {
		return model.Bag{}, err
	}
	return bag, nil
}
