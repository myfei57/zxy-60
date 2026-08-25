package checkin

import "bagsort/internal/model"

func (d *Desk) CancelCheckIn(bagID string) error {
	bag, ok, err := d.book.GetBag(bagID)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}
	bag.State = model.BagDiverted
	return d.book.PutBag(bag)
}
