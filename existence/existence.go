package existence

import (
	"GoIdeal/coupling"
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

type Existence struct {

	// explicitly state that the embedding struct needs to satisfy
	// the embedded interface
	Existencer
	mood Mood
	experiments map[string]coupling.Experiment
	results map[string]coupling.Result
	selfSatisfactionCounter int
	previousExperiment coupling.Experiment
}

// simulation of a stream of intelligence when it is run step by step
type Existencer interface {

	// one step of a  "stream of intelligence"
	// It returns a string representing an "event of intelligence" that was performed
	Step() string
}

