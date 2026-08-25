package flight

func (b *Book) Deadline(flightID string) (string, error) {
	f, err := b.Get(flightID)
	if err != nil {
		return "", err
	}
	return f.CutoffAt, nil
}

func (b *Book) SetCutoff(flightID string, cutoff string) error {
	snap, err := b.store.Load()
	if err != nil {
		return err
	}
	f, ok := snap.Flights[flightID]
	if !ok {
		return ErrNotFound
	}
	f.CutoffAt = cutoff
	snap.Flights[flightID] = f
	return b.store.Save(snap)
}
