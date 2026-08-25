package audit

import "bagsort/internal/model"

func (l *Logger) Count() (int, error) {
	records, err := l.List()
	if err != nil {
		return 0, err
	}
	return len(records), nil
}

func (l *Logger) ForBag(bagID string) ([]model.SortRecord, error) {
	records, err := l.List()
	if err != nil {
		return nil, err
	}
	out := make([]model.SortRecord, 0)
	for _, record := range records {
		if record.BagID == bagID {
			out = append(out, record)
		}
	}
	return out, nil
}
