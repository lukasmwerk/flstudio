package arranger

import "github.com/lukasmwerk/flstudio/pkg/wrapper"

// An instrument is a "tool" capable of generating offline musical elements.
// An offline musical element is one that we explicitly save, whereas online
// musical elements are processed live in the mixer.
// Instruments are loaded in via the FL wrapper library.

type Instrument struct {
	Name string

	Wrapper wrapper.Wrapper
}
