package sorter

func (s *Sorter) Verdict(epoch uint64, seq uint64) bool {
	return seq > s.lastSeq
}

func (s *Sorter) MarkSeen(epoch uint64, seq uint64) {
	s.lastEpoch = epoch
	s.lastSeq = seq
}
