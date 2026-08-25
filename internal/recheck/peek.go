package recheck

import "bagsort/internal/model"

func (q *Queue) Peek() (model.Bag, bool) {
	if len(q.items) == 0 {
		return model.Bag{}, false
	}
	return q.items[0], true
}

func (q *Queue) Items() []model.Bag {
	return q.items
}
