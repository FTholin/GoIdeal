package interaction

import (
	"GoIdeal/coupling"
	"fmt"
)

type Interaction010 struct {
	Label string
	coupling.Experiment
	coupling.Result
}


func (i Interaction010) String() string {
	return fmt.Sprintf("%s%s", i.Label, i.Result.Label)
}