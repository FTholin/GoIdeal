package interaction

import (
	"fmt"
)

type Interaction010 struct {
	Interaction000
}


func NewInteraction010(label string) *Interaction010 {

	interaction := &Interaction010{
		Interaction000: Interaction000{
			Label:label,
		},
	}
	return interaction
}

func (i Interaction010) String() string {
	return fmt.Sprintf("%s%s", i.Label, i.Result.Label)
}