package main

import (
	"fmt"
	"net/http"
	"time"
)

type maxLimiter struct {
	duration           time.Duration
	maxAllowedRequests int
	eventTimes         []time.Time
	count              counter
}

func NewMaxLimiter(duration time.Duration, maxAllowedRequests int) *maxLimiter {
	return &maxLimiter{
		duration:           duration,
		maxAllowedRequests: maxAllowedRequests,
		eventTimes:         []time.Time{},
		count:              counter{0},
	}
}

func (m *maxLimiter) Execute(req func() (*http.Response, error)) (interface{}, error) {

}

func main() {
	fmt.Println("hi")
}
