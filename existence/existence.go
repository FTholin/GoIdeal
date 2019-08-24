package existence

import (
	"GoIdeal/coupling"
	"GoIdeal/coupling/interaction"
)

type Mood int

const (
	SELF_SATISFIED = iota
	FRUSTRATED
	BORED
	PAINED
	PLEASED
)

func (mood Mood) String() string {
	values := [...]string{
		"SELF_SATISFIED",
		"FRUSTRATED",
		"BORED",
		"PAINED",
		"PLEASED"}

	if mood < SELF_SATISFIED || mood > PLEASED {
		return "Unknown"
	}
	return values[mood]
}

const LABEL_E1 = "e1"
const LABEL_E2 = "e2"
const LABEL_R1 = "r1"
const LABEL_R2 = "r2"

const BOREDOME_LEVEL = 4


type Existence interface {
	Step() string
	addOrGetPrimitiveInteraction(coupling.Experiment, coupling.Result) interaction.Interaction
	addOrGetInteraction(string) interaction.Interaction
	createInteraction(string) interaction.Interaction
	getInteraction(string) interaction.Interaction
	predict(coupling.Experiment) coupling.Result
	getOtherExperiment(coupling.Experiment) coupling.Experiment
}



