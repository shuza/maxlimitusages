package main

import "fmt"

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
