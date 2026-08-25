package flight

func (b *Book) AssignBag(bagID string, flightID string) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	snap.FlightMappings[bagID] = flightID
	return b.store.Save(snap)
}

func (b *Book) BagFlight(bagID string) (string, error) {
	snap, err := b.store.Load()
	if err != nil {
		return "", err
	}
	flightID, ok := snap.FlightMappings[bagID]
	if !ok {
		return "", ErrNotFound
	}
	return flightID, nil
}

func (b *Book) ScheduleChange(bagID string, flightID string) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	snap.FlightMappings[bagID] = flightID
	return b.store.Save(snap)
}

func (b *Book) SetCarousel(flightID string, carousel string) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	f, ok := snap.Flights[flightID]
	if !ok {
		return ErrNotFound
	}
	f.Carousel = carousel
	snap.Flights[flightID] = f
	return b.store.Save(snap)
}
