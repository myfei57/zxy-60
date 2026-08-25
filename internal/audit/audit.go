package audit

import (
	"bagsort/internal/model"
	"bagsort/internal/store"
)

type Logger struct {
	store *store.Store
}

func NewLogger(st *store.Store) *Logger {
	return &Logger{store: st}
}

func (l *Logger) Record(entry model.SortRecord) error {
	return l.store.AddSortRecord(entry)
}

func (l *Logger) List() ([]model.SortRecord, error) {
	snap, err := l.store.Load()
	if err != nil {
		return nil, err
	}
	return snap.SortRecords, nil
}
