package arranger

import "fmt"

// A Track is a group of clips in an arrangment,
// that are organized into different subtracks.
type Track struct {
	// A track has a unique name
	Name string

	// Every track has a set of instruments associated with it.
	// Instruments don't span tracks, that's where subtracks come in.
	// A track can have multiple instruments, for example for a drumset.
	Instruments []string

	// This flag tracks whether a track is expanded or not.
	expanded bool

	// This flag tracks whether a track should be rendered or not.
	render bool

	// A track doesn't hold any clips directly, that is done by subtracks.
	// Instead, the track when collapsed shows a minimap.
	// Todo minimap needs to be a type
	minimap string

	// A track keeps a reference to all its subtracks.
	// If a track is expanded, we render all the subtracks for the track.
	// Their order doesn't matter visually or on the backend and they can be reordered.
	subtracks []*Subtrack
}

func (t *Track) Render() error {
	if !t.render {
		return fmt.Errorf("render is set to false")
	}
	if t.expanded {
		return nil
	} else {
		return nil
	}
}
