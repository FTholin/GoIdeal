package interaction

import "GoIdeal/coupling"

type Interaction000 struct {
	Label string
	coupling.Experiment
	coupling.Result
}


func NewInteraction000(label string) *Interaction000 {

	interaction := &Interaction000{
		Label:label,
	}
	return interaction
}