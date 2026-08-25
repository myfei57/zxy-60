package checkin

import "bagsort/internal/model"

func (d *Desk) CheckIn(barcode string, flightID string) (model.Bag, error) {
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
