package counters

type alertCounter int // unexported

// New creates and returns values of the unexported type alertCounter
func New(value int) alertCounter {
	return alertCounter(value)
}
