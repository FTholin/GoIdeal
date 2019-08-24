package existence

import (
	"GoIdeal/coupling"
	"GoIdeal/coupling/interaction"
	"fmt"
)


// An Existence010 simulates a "stream of intelligence" made of a succession of Experiments and Results.
// Existence010 is SELF-SATISFIED when the Result corresponds to the Result it expected, and FRUSTRATED otherwise.
// Additionally, the Existence0 is BORED when it has been SELF-SATISFIED for too long, which causes it to try another Experiment.
// An Existence1 is still a single entity rather than being split into an explicit Agent and Environment.
type Existence010 struct {
	Existence000
	interactions map[string]interaction.Interaction010
}


func NewExistence010() *Existence010 {

	experiments := map[string]coupling.Experiment{
		LABEL_E1: coupling.Experiment{LABEL_E1},
		LABEL_E2: coupling.Experiment{LABEL_E2},
	}
	existence := &Existence010{
		Existence000: Existence000{
			experiments: experiments,
			results: make(map[string]coupling.Result),
			selfSatisfactionCounter: 0,
			previousExperiment: experiments[LABEL_E1],
		},
		interactions: make(map[string]interaction.Interaction010),
	}
	return existence
}

func (e *Existence010) Step() string {
	experiment := e.previousExperiment

	if e.mood == BORED {
		experiment = e.getOtherExperiment(experiment)
		e.selfSatisfactionCounter = 0
	}

	anticipatedResult := e.predict(experiment)
	result := e.ReturnResult010(experiment)

	e.addOrGetPrimitiveInteraction(experiment, result)

	if result == anticipatedResult {
		e.mood = SELF_SATISFIED
		e.selfSatisfactionCounter += 1
	} else {
		e.mood = FRUSTRATED
		e.selfSatisfactionCounter = 0
	}

	if e.selfSatisfactionCounter >= BOREDOME_LEVEL {
		e.mood = BORED
	}

	e.previousExperiment = experiment

	return fmt.Sprintf("%s%s %s", experiment.Label, result.Label, e.mood)
}

// Create an interaction as a tuple <experiment, result>
func (e *Existence010) addOrGetPrimitiveInteraction(experiment coupling.Experiment, result coupling.Result) {
	interaction := e.addOrGetInteraction(fmt.Sprintf("%s%s", experiment.Label, result.Label))
	interaction.Experiment = experiment
	interaction.Result = result
	e.interactions[interaction.Label] = interaction
}


// records an interaction in memory.
func (e *Existence010) addOrGetInteraction(label string) interaction.Interaction010 {
	if _, contains := e.interactions[label]; !contains {
		interaction := interaction.NewInteraction010(label)
		e.interactions[label] = *interaction
	}
	return e.interactions[label]
}

// Finds an interaction from its experiment
func (e *Existence010) predict(experiment coupling.Experiment) coupling.Result {
	var anticipatedResult coupling.Result
	for _, v := range e.interactions {
		if v.Experiment == experiment {
			anticipatedResult = v.Result
		}
	}
	return anticipatedResult
}

// Creates a new experiment from its label and stores it in memory.
func (e *Existence010) addOrGetExperiment(label string) coupling.Experiment {
	if _, contains := e.experiments[label]; !contains {
		e.experiments[label] = coupling.Experiment{
			Label: label,
		}
	}
	return e.experiments[label]
}

//// Finds an experiment different from that passed in parameter.
//func (e *Existence010) getOtherExperiment(experiment coupling.Experiment) coupling.Experiment {
//	var otherExperiment coupling.Experiment
//
//	for _, v := range e.interactions {
//		if v.Experiment != experiment {
//			otherExperiment = v.Experiment
//			break
//		}
//	}
//	return otherExperiment
//}

// Creates a new result from its label and stores it in memory.
//func (e *Existence010) createOrGetResult(label string) coupling.Result {
//	if _, contains := e.results[label]; !contains {
//		e.results[label] = coupling.Result{
//			Label: label,
//		}
//	}
//	return e.results[label]
//}

//// Environement010
//// E1 results in R1. E2 result in R2.
//func (e *Existence010) ReturnResult010(experiment coupling.Experiment) coupling.Result {
//	if experiment == e.addOrGetExperiment(LABEL_E1) {
//		return e.createOrGetResult(LABEL_R1)
//	} else {
//		return e.createOrGetResult(LABEL_R2)
//	}
//}