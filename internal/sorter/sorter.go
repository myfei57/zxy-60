package sorter

import (
	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/store"
)

type Sorter struct {
	store     *store.Store
	book      *flight.Book
	chutes    *chute.Assigner
	lastEpoch uint64
	lastSeq   uint64
}

func NewSorter(st *store.Store, book *flight.Book, chutes *chute.Assigner) *Sorter {
	return &Sorter{
		store:  st,
		book:   book,
		chutes: chutes,
	}
}
