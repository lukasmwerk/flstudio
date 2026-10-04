// Monitor is a plugin we might just integrate into the DAW. Or make it so that plugins can
// achieve this integration cleanly.
// Monitor has the following main capabilities:
// Loudness measurements
// A|B listening (if not integrated into DAW)
// Frequency spectrum analyzer
// Soundstage analyzer
// Levels analyzer
// Notes analyzer (harmonies/chords/intonation)
// Masking analyzer & muddiness, etc.
// Transients analyzer
// Monitor is not an all in one mastering plugin. It's a master/bus channel
// analyzing plugin, where all the analyzers are packaged into a cohesive,
// intuitive unit rather than a bunch of separate UIs and tabs.
// The intended use is to load monitor at the end of a patcher chain,
// usually on a monitor channel, usually by default, and use it to gain
// extra insights about the sound on that channel
package monitor
