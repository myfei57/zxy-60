package store

import "bagsort/internal/model"

func (s *Store) AppendTransferCommand(cmd model.TransferCommand) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	snap.TransferCommands = append(snap.TransferCommands, cmd)
	return s.Save(snap)
}

func (s *Store) TransferCommands() ([]model.TransferCommand, error) {
	snap, err := s.Load()
	if err != nil {
		return nil, err
	}
	return snap.TransferCommands, nil
}

func (s *Store) ClearTransferCommandsBefore(batch uint64) error {
	snap, err := s.Load()
	if err != nil {
		return err
	}
	kept := make([]model.TransferCommand, 0, len(snap.TransferCommands))
	for _, cmd := range snap.TransferCommands {
		if cmd.Batch >= batch {
			kept = append(kept, cmd)
		}
	}
	snap.TransferCommands = kept
	return s.Save(snap)
}
