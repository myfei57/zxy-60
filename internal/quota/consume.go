package quota

func (c *Checker) Consume(flightID string, n int) bool {
	limit := c.Limit(flightID)
	if limit == 0 {
		return true
	}
	c.limits[flightID] = limit - n
	return c.limits[flightID] >= 0
}

func (c *Checker) Release(flightID string, n int) {
	c.limits[flightID] = c.Limit(flightID) + n
}

func (c *Checker) Limits() map[string]int {
	out := map[string]int{}
	for flightID, limit := range c.limits {
		out[flightID] = limit
	}
	return out
}
