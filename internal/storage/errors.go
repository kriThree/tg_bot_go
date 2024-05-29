package storage

type StorageError struct {
	text string
}

func (s StorageError) Error() string {
	return s.text
}

func new(text string) StorageError {
	return StorageError{text: text}
}

var (
	MeaningAlreadyDefined = new("meaning already defined")
	InternalError         = new("internal server error")
	DefinitionNotFound    = new("definition not found")
	MeaningNotFound       = new("meaning not found")
)
