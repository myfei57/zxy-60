package quota

type Checker struct {
	limits map[string]int
}

func NewChecker() *Checker {
	return &Checker{limits: map[string]int{}}
}

func (c *Checker) SetLimit(flightID string, limit int) {
	c.limits[flightID] = limit
}

func (c *Checker) Limit(flightID string) int {
	if v, ok := c.limits[flightID]; ok {
		return v
	}
	return 0
}

func (c *Checker) Allowed(flightID string, current int) bool {
	limit := c.Limit(flightID)
	return limit == 0 || current < limit
}
