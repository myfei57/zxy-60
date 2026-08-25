package audit

func (l *Logger) FlightSummary(flightID string) (map[string]int, error) {
	records, err := l.List()
	if err != nil {
		return nil, err
	}
	summary := map[string]int{}
	for _, record := range records {
		if record.FlightID == flightID {
			summary[record.ChuteID]++
		}
	}
	return summary, nil
}
