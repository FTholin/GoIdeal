package existence

// simulation of a stream of intelligence when it is run step by step
type Existence interface {

	// one step of a  "stream of intelligence"
	// It returns a string representing an "event of intelligence" that was performed
	Step() string
}

