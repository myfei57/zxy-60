package tag

import "bagsort/internal/model"

func (r *Reader) LastSequence() (model.Sequence, error) {
	snap, err := r.store.Load()
	if err != nil {
		return model.Sequence{}, err
	}
	return snap.Sequence, nil
}

func (r *Reader) ReadBatch(barcodes []string) ([]Reading, error) {
	out := make([]Reading, 0, len(barcodes))
	for _, barcode := range barcodes {
		reading, err := r.Read(barcode)
		if err != nil {
			return nil, err
		}
		out = append(out, reading)
	}
	return out, nil
}
