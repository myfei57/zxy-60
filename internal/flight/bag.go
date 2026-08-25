package flight

import "bagsort/internal/model"

func (b *Book) GetBag(bagID string) (model.Bag, bool, error) {
	return b.store.GetBag(bagID)
}

func (b *Book) PutBag(bag model.Bag) error {
	return b.store.PutBag(bag)
}

func (b *Book) ListBags() ([]model.Bag, error) {
	return b.store.ListBags()
}
