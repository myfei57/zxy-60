package tag

import "bagsort/internal/model"

const MaxSequence uint64 = 1000

func Allocate(seq model.Sequence) model.Sequence {
	seq.Next++
	if seq.Next > MaxSequence {
		seq.Epoch++
		seq.Next = 1
	}
	return seq
}
