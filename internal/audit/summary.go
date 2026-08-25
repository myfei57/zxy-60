package audit

import "bagsort/internal/model"

func (l *Logger) Last() (model.SortRecord, bool, error) {
	records, err := l.List()
	if err != nil {
		return model.SortRecord{}, false, err
	}
	if len(records) == 0 {
		return model.SortRecord{}, false, nil
	}
	return records[len(records)-1], true, nil
}

func (l *Logger) Summary() (map[string]int, error) {
	records, err := l.List()
	if err != nil {
		return nil, err
	}
	summary := map[string]int{}
	for _, record := range records {
		summary[record.ChuteID]++
	}
	return summary, nil
}
