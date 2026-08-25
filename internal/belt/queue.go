package belt

import (
	"time"

	"bagsort/internal/model"
)

func (b *Belt) AppendTransfer(flightID string, carousel string, batch uint64) error {
	cmd := model.TransferCommand{
		ID:        newCommandID(),
		FlightID:  flightID,
		Carousel:  carousel,
		Batch:     batch,
		CreatedAt: time.Now().UTC().Format(time.RFC3339),
	}
	return b.store.AppendTransferCommand(cmd)
}

func (b *Belt) TransferCommands() ([]model.TransferCommand, error) {
	return b.store.TransferCommands()
}

func (b *Belt) ReplayTransfers() ([]model.TransferCommand, error) {
	cmds, err := b.store.TransferCommands()
	if err != nil {
		return nil, err
	}
	out := make([]model.TransferCommand, 0, len(cmds))
	for _, cmd := range cmds {
		if cmd.Batch == b.batch {
			out = append(out, cmd)
		}
	}
	return out, nil
}

func (b *Belt) discardStaleCommands(nextBatch uint64) error {
	return b.store.ClearTransferCommandsBefore(nextBatch)
}

func newCommandID() string {
	return "cmd-" + time.Now().UTC().Format("20060102150405.000000000")
}
