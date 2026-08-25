package checkin

import "bagsort/internal/model"

type Entry struct {
	Barcode  string `json:"barcode"`
	FlightID string `json:"flight_id"`
}

func (d *Desk) FindBag(bagID string) (model.Bag, bool, error) {
	return d.book.GetBag(bagID)
}

func (d *Desk) CheckInBatch(entries []Entry) ([]model.Bag, error) {
	out := make([]model.Bag, 0, len(entries))
	for _, entry := range entries {
		bag, err := d.CheckIn(entry.Barcode, entry.FlightID)
		if err != nil {
			return nil, err
		}
		out = append(out, bag)
	}
	return out, nil
}
