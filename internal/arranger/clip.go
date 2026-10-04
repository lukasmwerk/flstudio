package arranger

// A clip is a virtual slice/window of audio.
// The audio itself can be generated and optimized by the mixer at runtime.
// Audio can come from an "instrument" or a "soundfile".
type Clip struct {
}

func NewClip() (*Clip, error) {
	return nil, nil
}

// Clips are saved to the arrangement cliplist (injected?)
// TODO decided where this should be done from
func (c *Clip) Save(r Cliplist) error {
	return nil
}
