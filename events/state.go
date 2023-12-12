package events

import (
	"github.com/dotabook/go-dota2/state"
)

// ClientStateChanged is emitted whenever anything about the client state changes.
type ClientStateChanged struct {
	// OldState is the old state.
	OldState state.Dota2State
	// NewState is the new state.
	NewState state.Dota2State
}
