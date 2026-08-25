package belt

import "bagsort/internal/model"

// Switch moves the belt to a new carousel and batch. The current batch's bags
// must be cleared BEFORE the carousel/batch fields are updated: clearCurrent
// attributes leftover (unclaimed) bags to b.carousel/b.batch, so those fields
// must still hold the OLD values when it runs. Clearing first means leftover
// bags stay attributed to their original carousel, and the new carousel starts
// empty — no old-batch residue is carried onto the next batch.
func (b *Belt) Switch(nextCarousel string, nextBatch uint64) error {
	if err := b.clearCurrent(); err != nil {
		return err
	}
	b.carousel = nextCarousel
	b.batch = nextBatch
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
