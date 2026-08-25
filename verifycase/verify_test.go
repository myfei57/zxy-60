package verifycase

import (
	"testing"

	"bagsort/internal/chute"
	"bagsort/internal/flight"
	"bagsort/internal/sorter"
	"bagsort/internal/store"
	"bagsort/internal/tag"
)

func TestBsTagEpochRollover(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	chutes := chute.NewAssigner(st)
	srt := sorter.NewSorter(st, book, chutes)

	srt.MarkSeen(0, tag.MaxSequence)
	if !srt.Verdict(1, 1) {
		t.Fatal("expected fresh bag after epoch rollover to be accepted")
	}
}
