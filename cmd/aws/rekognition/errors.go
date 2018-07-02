package rekognition

const (

	// ErrEmptyBytes is used when an empty []byte is passed to a method
	// that requires a full one
	ErrEmptyBytes = "EmptyBytes"

	// ErrBadSimilarity is used when 0 is passed as similarity parameter
	ErrBadSimilarity = "BadSimilarity"

	// ErrEmptyParameter is used when a parameter is empty
	ErrEmptyParameter = "EmptyParameter"

	// ErrEmptyMap is used when structs.Map() returns an empty map
	ErrEmptyMap = "EmptyMap"
)
