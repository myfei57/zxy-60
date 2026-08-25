package tag

func (r *Reader) CommitBatch(readings []Reading) error {
	for _, reading := range readings {
		if err := r.Commit(reading); err != nil {
			return err
		}
	}
	return nil
}
