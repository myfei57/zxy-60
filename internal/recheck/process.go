package recheck

import (
	"bagsort/internal/flight"
	"bagsort/internal/model"
)

type Processor struct {
	queue *Queue
	book  *flight.Book
}

func NewProcessor(q *Queue, book *flight.Book) *Processor {
	return &Processor{queue: q, book: book}
}

func (p *Processor) Next() (model.Bag, bool) {
	return p.queue.Dequeue()
}

func (p *Processor) Deadline(flightID string) (string, error) {
	return p.book.Deadline(flightID)
}
