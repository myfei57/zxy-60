package sorter

func (s *Sorter) AssignChute(flightID string, chuteID string) error {
	return s.chutes.Assign(flightID, chuteID)
}

func (s *Sorter) CurrentChute(flightID string) (string, error) {
	return s.chutes.Current(flightID)
}
