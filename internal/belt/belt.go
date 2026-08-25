package belt

import (
	"bagsort/internal/flight"
	"bagsort/internal/model"
	"bagsort/internal/store"
)

type Belt struct {
	store    *store.Store
	book     *flight.Book
	carousel string
	batch    uint64
	onCarousel []model.Bag
	loads    []model.LoadRecord
}

func NewBelt(st *store.Store, book *flight.Book) *Belt {
	return &Belt{
		store:      st,
		book:       book,
		onCarousel: []model.Bag{},
		loads:      []model.LoadRecord{},
	}
}

func (b *Belt) Carousel() string {
	return b.carousel
}

func (b *Belt) Batch() uint64 {
	return b.batch
}

func (b *Belt) Place(bag model.Bag) {
	b.onCarousel = append(b.onCarousel, bag)
}

func (b *Belt) OnCarousel() []model.Bag {
	return b.onCarousel
}

func (b *Belt) Loads() []model.LoadRecord {
	return b.loads
}
