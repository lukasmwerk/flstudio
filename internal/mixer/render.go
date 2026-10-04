package mixer

// A render is a rendered snippet of a recording. This isn't really
// surfaced to the user except when rendering/saving recordings to files.
// Renders are highly optimized to run in the background during normal
// operations, to reduce CPU load on playback.
type Render struct {
}
