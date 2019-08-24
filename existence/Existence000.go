package existence

import (
	"GoIdeal/coupling"
	"GoIdeal/coupling/interaction"
)

type Existence000 struct {
	mood Mood
	experiments map[string]coupling.Experiment
	results map[string]coupling.Result
	interactions map[string]interaction.Interaction000
	selfSatisfactionCounter int
	previousExperiment coupling.Experiment
}

// Finds an experiment different from that passed in parameter.
func (e *Existence000) getOtherExperiment(experiment coupling.Experiment) coupling.Experiment {
	var otherExperiment coupling.Experiment

	for _, v := range e.interactions {
		if v.Experiment != experiment {
			otherExperiment = v.Experiment
			break
		}
	}
	return otherExperiment
}


// Creates a new result from its label and stores it in memory.
func (e *Existence000) createOrGetResult(label string) coupling.Result {
	if _, contains := e.results[label]; !contains {
		e.results[label] = coupling.Result{
			Label: label,
		}
	}
	return e.results[label]
}

// Create an interaction as a tuple <experiment, result>
//func (e *Existence000) addOrGetPrimitiveInteraction(experiment coupling.Experiment, result coupling.Result) {
//	interaction := e.addOrGetInteraction(fmt.Sprintf("%s%s", experiment.Label, result.Label))
//	interaction.Experiment = experiment
//	interaction.Result = result
//	e.interactions[interaction.Label] = interaction
//}


//// records an interaction in memory.
//func (e *Existence000) addOrGetInteraction(label string) interaction.Interaction000 {
//	if _, contains := e.interactions[label]; !contains {
//		interaction := interaction.NewInteraction000(label)
//		e.interactions[label] = *interaction
//	}
//	return e.interactions[label]
//}

// Creates a new experiment from its label and stores it in memory.
func (e *Existence000) addOrGetExperiment(label string) coupling.Experiment {
	if _, contains := e.experiments[label]; !contains {
		e.experiments[label] = coupling.Experiment{
			Label: label,
		}
	}
	return e.experiments[label]
}

// Environement010
// E1 results in R1. E2 result in R2.
func (e *Existence000) ReturnResult010(experiment coupling.Experiment) coupling.Result {
	if experiment == e.addOrGetExperiment(LABEL_E1) {
		return e.createOrGetResult(LABEL_R1)
	} else {
		return e.createOrGetResult(LABEL_R2)
	}
}

