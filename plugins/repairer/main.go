// Repairer is a plugin for visually identifying artifacts and removing them from audio.
// It typically goes on a tools aux channel routed directly to a monitor. That way, it doesn't
// interfere with the rest of the mix, and users can specifically jump into the tool to remove
// artifacts for a recording. Ideally, we can build the mixer such that the channels have access
// to the offline renders + recordings. That way the repairer can scrub through the whole recording
// selected.
package repairer
