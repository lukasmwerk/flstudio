package mixer

// The Mixer is the subsystem that allows for audio to be routed and processed.
type Mixer struct {
}

// The mixer is constantly rendering recordings through the online processing.
// Recordings aren't a format that a speaker can play, so we have to render it to
// a raw format a speaker can play.
// The DAW tracks where it can safely cache renders to improve performance.
// Most playback is of elements in the mix tree that are unchanged since the last play.
// Parts of the mix tree are rendered, while others are streamed through live.
// That way, playback is optimized, but the user must never notice.
// We use

// We can hear/feel
