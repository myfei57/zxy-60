package checkin

import (
	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/inject"
	"bagsort/internal/quota"
	"bagsort/internal/tag"
)

type Desk struct {
	reader   *tag.Reader
	book     *flight.Book
	chutes   *chute.Assigner
	injector *inject.Injector
	quota    *quota.Checker
}

func NewDesk(reader *tag.Reader, book *flight.Book, chutes *chute.Assigner, injector *inject.Injector, quota *quota.Checker) *Desk {
	return &Desk{
		reader:   reader,
		book:     book,
		chutes:   chutes,
		injector: injector,
		quota:    quota,
	}
}
