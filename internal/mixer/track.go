package mixer

// A Track is a mixer channel that corresponds to an Arrangement track.
// A track has a patcher (we replace FL's track slots with a single patcher instance).
// This way its equally easy to create both simple and complex effect chains.
// The patcher opens as a pop up for editing in the arranger pane.
// The mixer pane shows only a simple summary/minimap of the chain.
type Track struct {
}


// A track is gainstaged. This is a digital format with no benefit of redlining
// the channel. All tracks are gainstaged based on the calculated max. The max 
// can only be topped during live recording.
