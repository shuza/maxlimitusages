package main

import (
	"fmt"
	"time"
)

type maxLimiter struct {
	duration           time.Duration
	maxAllowedRequests int
	eventTimes         []time.Time
}

func NewMaxLimiter(duration time.Duration, maxAllowedRequests int) *maxLimiter {
	return &maxLimiter{
		duration:           duration,
		maxAllowedRequests: maxAllowedRequests,
		eventTimes:         []time.Time{},
	}
}

// State is a type that represents a state of the limit
type State int

const (
	StateDenied State = iota
	StateAllow
	StateWait
	StateMaxRequest
	StateTooManyRequest
)

// String implements stringer interface
func (s State) String() string {
	switch s {
	case StateDenied:
		return "Denied"
	case StateAllow:
		return "Allow"
	case StateWait:
		return "Wait"
	case StateMaxRequest:
		return "MaxRequest"
	case StateTooManyRequest:
		return "TooManyRequest"
	default:
		return fmt.Sprintf("unknown state: %d", s)
	}
}

type Counts struct {
	ConsecutiveRequests uint32
}

func (c *Counts) onRequest() {
	c.ConsecutiveRequests++
}

func (c *Counts) clear() {
	c.ConsecutiveRequests = 0
}

func main() {
	fmt.Println("hi")
}
