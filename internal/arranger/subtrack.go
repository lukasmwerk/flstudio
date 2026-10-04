package arranger

// A subtrack is a child of a track,
// it allows a track to stick to a single instrument
// in most situations, while allowing clips to be arranged
// by performance or arbitrary visual hierarchy.
type Subtrack struct {
	Clips
}
