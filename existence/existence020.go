package existence

import (
	"GoIdeal/coupling"
	"GoIdeal/coupling/interaction"
)

type Existence020 struct {
	Existence
	interactions map[string]interaction.Interaction020
}

func NewExistence020() *Existence020 {
	ex := &Existence020{

	}
	return ex
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

	return interaction.Interaction020{
		Interaction010: interaction.Interaction010{
			Label: label,
		},
	}
}