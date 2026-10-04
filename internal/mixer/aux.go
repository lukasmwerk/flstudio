package mixer

type channel interface {
}

// An Aux is an auxiliary channel in the Mixer.
// It allows you to route audio from any channel(s),
// and back to any channel(s).
type Aux struct {
}
