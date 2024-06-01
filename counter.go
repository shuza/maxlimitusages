package main

type counter struct {
	consecutiveRequests uint32
}

func (c *counter) onRequest() {
	c.consecutiveRequests++
}

func (c *counter) clear() {
	c.consecutiveRequests = 0
}
