package valueObjects

import (
	"errors"
)

const StatusCodeBadActionError = 499

var (
	ErrInvalidLanguage                                   = errors.New("INVALID_LANGUAGE")
)