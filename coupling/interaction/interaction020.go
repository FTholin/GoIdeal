package interaction

import "fmt"

// An Interaction020 is an Interaction010 with a valence.
type Interaction020 struct {
	Interaction010
	Valence int
}


func (i Interaction020) String() string {
	return fmt.Sprintf("%s, %d", i.Label, i.Valence)
}