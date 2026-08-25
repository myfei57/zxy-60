package recheck

import "bagsort/internal/model"

type Queue struct {
	items []model.Bag
}

func NewQueue() *Queue {
	return &Queue{items: []model.Bag{}}
}

func (q *Queue) Enqueue(bag model.Bag) {
	q.items = append(q.items, bag)
}

func (q *Queue) Dequeue() (model.Bag, bool) {
	if len(q.items) == 0 {
		return model.Bag{}, false
	}
	first := q.items[0]
	q.items = q.items[1:]
	return first, true
}

func (q *Queue) Len() int {
	return len(q.items)
}
