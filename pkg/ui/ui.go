// UI is the library responsible for connecting fl to the user.
// This includes the graphical interface, tty, mouse, midi, consoles, etc.
// Ultimately, the arranger, explorer and mixer handle their own user interface
// business logic, so UI is just a wrapper for OS and device APIs, and an abstraction
// layer that reduces the interface area for the business logic to something useful.
// The philosophy behind this codebase is to couple whats happening in the UI more with
// what's happening on the backend.
// The resources are similar to graphical tools, so its easy to map concepts from designs in Figma, etc.
package ui
