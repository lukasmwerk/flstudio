package mixer

// A recording is a collection of audio data that has been recorded together.
// Internally, we are able to store and split recordings by harmony, instruments, performers, etc.
// This allows us to do things like move around instruments in the space they were recorded in,
// change the harmonies, rhythms, performance nuances, performers, inspirations, space details, etc.
// We can split out elements, insert elements from different recordings, and give this interface
// to plugins for creative use-cases.
// Naturally because we have to make heavy use of ML for this we can't expect full determinism and reproducability,
// so recordings are saved versioned.
// Clips in the playlist window recordings (or renders at runtime), the latter two living in the mixer's runtime.
// Recordings are both offline and online and can be used for both midi instruments and audio recordings -- which are the same in our DAW.
type Recording struct {
}

// How to handle midi notes and broken down (split) harmony sounds.
