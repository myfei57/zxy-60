package belt

import "bagsort/internal/model"

func (b *Belt) SetCarousel(carousel string, batch uint64) {
	b.carousel = carousel
	b.batch = batch
}

func (b *Belt) CurrentBatch() uint64 {
	return b.batch
}

func (b *Belt) PlaceBatch(bags []model.Bag) {
	b.onCarousel = append(b.onCarousel, bags...)
}
