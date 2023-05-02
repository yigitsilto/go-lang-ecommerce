package exceptions

import "errors"

var (
	EntityNotFoundException = errors.New("Entity not found")

	DuplicateValueException = errors.New("This data is already exists")

	ServerError = errors.New("There has been an error. Please try again later")
)
