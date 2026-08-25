package belt

func (b *Belt) Rotate() error {
	if b.carousel == "" {
		return nil
	}
	return b.discardStaleCommands(b.batch)
}
