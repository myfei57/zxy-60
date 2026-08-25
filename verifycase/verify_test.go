package verifycase

import (
	"testing"

	"bagsort/internal/belt"
	"bagsort/internal/flight"
	"bagsort/internal/store"
)

func TestBsNoTransferReplay(t *testing.T) {
	st := store.New(t.TempDir(), "bagsort.json")
	book := flight.NewBook(st)
	b := belt.NewBelt(st, book)

	if err := b.AppendTransfer("F1", "A", 1); err != nil {
		t.Fatal(err)
	}
	if err := b.AppendTransfer("F1", "B", 2); err != nil {
		t.Fatal(err)
	}
	if err := b.Switch("B", 2); err != nil {
		t.Fatal(err)
	}

	cmds, err := b.TransferCommands()
	if err != nil {
		t.Fatal(err)
	}
	if len(cmds) != 1 {
		t.Fatalf("expected only batch 2 command to remain, got %d", len(cmds))
	}
}
