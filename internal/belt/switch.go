package belt

import "bagsort/internal/model"

func (b *Belt) Switch(nextCarousel string, nextBatch uint64) error {
	b.carousel = nextCarousel
	b.batch = nextBatch
	if err := b.clearCurrent(); err != nil {
		return err
	}
	return b.discardStaleCommands(nextBatch)
}

func (b *Belt) clearCurrent() error {
	for _, bag := range b.onCarousel {
		b.loads = append(b.loads, loadRecord(bag.ID, b.carousel, b.batch))
	}
	b.onCarousel = nil
	return nil
}

func loadRecord(bagID string, carousel string, batch uint64) model.LoadRecord {
	return model.LoadRecord{BagID: bagID, Carousel: carousel, Batch: batch}
}
