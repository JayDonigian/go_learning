package cyoa

// CYOA is the text representation of the story. It is read in from a JSON file.
type CYOA struct {
	Name      string
	Title     string
	Adventure []string
	Options   []Option
}

// An Option is the representation of the end of a Story
type Option struct {
	Text string
	Arc  string
}

// New creates a CYOA based on a given JSON file.
func New() {

}
