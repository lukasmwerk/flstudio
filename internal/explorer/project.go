package explorer

// An FL project is a file and directory that saves the state of a
// session so that future sessions can start where we left off.
type Project struct {

	// Every project has a mutable name.
	Name string

	// The name under which the project would be published.
	Artist string

	// A map of collaborators to roles
	Credits map[string]string

	// Notes are associated with every project.
	// Maybe we use a file and/or buffer instead.
	Notes string

	// Project path: where the project folder is stored.
	Path string
}

func NewProject() (*Project, error) {
	return nil, nil
}

func LoadProject(path string) (*Project, error) {
	return nil, nil
}

func (p *Project) Save() error {
	return nil
}
