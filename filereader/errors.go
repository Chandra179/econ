package filereader

import (
	"errors"
	"fmt"
)

// Error definitions for the filereader package
var (
	ErrUnsupportedFileType = errors.New("unsupported file type")
	ErrReadingFile         = errors.New("error reading file")
	ErrParsingFile         = errors.New("error parsing file")
)

// WrapError wraps an original error with a custom message
func WrapError(err error, msg string) error {
	return fmt.Errorf("%s: %w", msg, err)
}
