package flight

func (b *Book) Batch(flightID string) (uint64, error) {
	f, err := b.Get(flightID)
	if err != nil {
		return 0, err
	}
	return f.Batch, nil
}

func (b *Book) SetBatch(flightID string, batch uint64) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	f, ok := snap.Flights[flightID]
	if !ok {
		return ErrNotFound
	}
	f.Batch = batch
	snap.Flights[flightID] = f
	return b.store.Save(snap)
}
