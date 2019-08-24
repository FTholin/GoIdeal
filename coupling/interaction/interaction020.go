package interaction

import "fmt"

// An Interaction020 is an Interaction010 with a valence.
type Interaction020 struct {
	Interaction000
	Valence int
}


func NewInteraction020(label string) *Interaction020 {

	interaction := &Interaction020{
		Interaction000: Interaction000{
			Label:label,
		},
	}
	return interaction
}


func (i Interaction020) String() string {
	return fmt.Sprintf("%s, %d", i.Label, i.Valence)
}