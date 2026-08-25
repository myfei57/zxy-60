package quota

func (c *Checker) Available(flightID string, current int) int {
	limit := c.Limit(flightID)
	if limit == 0 {
		return -1
	}
	if current >= limit {
		return 0
	}
	return limit - current
}
