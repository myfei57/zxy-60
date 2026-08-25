package chute

func (a *Assigner) Reassign(flightID string, chuteID string) error {
	return a.Assign(flightID, chuteID)
}
