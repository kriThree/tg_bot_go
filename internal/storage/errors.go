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
	MeaningAlreadyDefinedErr = new("meaning already defined")
	InternalErr              = new("internal server error")
	DefinitionNotFoundErr    = new("definition not found")
	MeaningNotFoundErr       = new("meaning not found")
	UserNotFoundErr          = new("user not found")
	UserAlreadyAddedErr      = new("user already added")
)
