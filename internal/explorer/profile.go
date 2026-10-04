package explorer

// A Profile is a configuration of FL or a project
type Profile struct {

	// An identifier for the profile
	Name string

	// A json compatible map of key value pairs that holds all the configurations
	// we should hardcode what configs are possible but maybe not here.
	Settings map[string]any
}
