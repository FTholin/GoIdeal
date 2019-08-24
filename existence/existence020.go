package existence

import (
	"GoIdeal/coupling"
	"GoIdeal/coupling/interaction"
	"fmt"
)

type Existence020 struct {
	existence Existence000
	interactions map[string]interaction.Interaction020
}

func NewExistence020() *Existence020 {

	experiments := map[string]coupling.Experiment{
		LABEL_E1: coupling.Experiment{LABEL_E1},
		LABEL_E2: coupling.Experiment{LABEL_E2},
	}

	results := map[string]coupling.Result{
		LABEL_R1: coupling.Result{LABEL_R1},
		LABEL_R2: coupling.Result{LABEL_R2},
	}

	existence := &Existence020{
		existence: Existence000{
			experiments: experiments,
			results: results,
			previousExperiment: experiments[LABEL_E1],
		},
		interactions: make (map[string]interaction.Interaction020),
	}

	/** Alter the valence of interactions to change the agent's motivation */
	existence.addOrGetPrimitiveInteraction(experiments[LABEL_E1],results[LABEL_R1],-1)
	existence.addOrGetPrimitiveInteraction(experiments[LABEL_E1],results[LABEL_R2],1)
	existence.addOrGetPrimitiveInteraction(experiments[LABEL_E2],results[LABEL_R1],-1)
	existence.addOrGetPrimitiveInteraction(experiments[LABEL_E2],results[LABEL_R2],1)

	return existence
}

func (e *Existence020) Step() string {
	experiment := e.existence.previousExperiment

	if e.existence.mood == PAINED {
		experiment = e.existence.getOtherExperiment(experiment)
	}

	result := e.existence.ReturnResult010(experiment)

	enactedInteraction := e.addOrGetPrimitiveInteraction(experiment, result, 0)

	if enactedInteraction.Valence >= 0 {
		e.existence.mood = PLEASED
	} else {
		e.existence.mood = PAINED
	}

	e.existence.previousExperiment = experiment

	return fmt.Sprintf("%s%s %s", experiment.Label, result.Label, e.existence.mood)
}


// Create an interaction as a tuple <experience, result>.
func (e *Existence020) addOrGetPrimitiveInteraction(experiment coupling.Experiment,
	result coupling.Result, valence int) interaction.Interaction020 {
	label := experiment.Label + result.Label

	if _, contains := e.interactions[label]; !contains {
		newInter := e.createInteraction(label)
		newInter.Experiment = experiment
		newInter.Result = result
		newInter.Valence = valence
		e.interactions[label] = newInter
	}
	return e.interactions[label]
}

func (e *Existence020) createInteraction(label string) interaction.Interaction020 {
	return *interaction.NewInteraction020(label)
}

func (e *Existence020) getInteraction(label string) interaction.Interaction020 {
	return e.interactions[label]
}


