package arranger

// An Arrangement is an organization of clips and tracks.
// Using multiple arrangements works like using branches.
// You can develop ideas in one arrangement and merge them to another arrangment.
// Arrangements can also be used for different distributions/mixes of the same song.
// Arrangements don't encode information about the playback or online processing of the audio.
// Arrangements also improve the experience for parallel collaboration on a project.
type Arrangement struct {
}

func NewArrangement() (*Arrangement, error) {
	return nil, nil
}

// Clone an arrangement
func (a *Arrangement) Clone() (*Arrangement, error) {
	return nil, nil
}

// Orchestrates a merge of an arrangement.
func (a *Arrangement) Merge() error {
	return nil
}

// Archives an arrangement file?
func (a Arrangement) Archive() error {
	return nil
}

// We save arrangements to files. Keeping with our principles,
// an arrangement doesn't save any information on the sound generation
// and playback, only how the musical elements are arranged.
func (a *Arrangement) Save() error {
	return nil
}
