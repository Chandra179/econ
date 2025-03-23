package filereader

import (
	"io"
)

// Reader defines the core interface for reading different file formats
type Reader interface {
	// Read parses the file at given path and returns the data as a slice of string slices
	// where each inner slice represents a row of data
	Read(filePath string) ([][]string, error)

	// ReadFromBytes parses file data from a byte slice
	ReadFromBytes(data []byte) ([][]string, error)

	// ReadFromReader parses file data from an io.Reader
	ReadFromReader(reader io.Reader) ([][]string, error)
}

// FileType represents supported file formats
type FileType string

const (
	CSV FileType = "csv"
	PDF FileType = "pdf"
)

// NewReader creates a new Reader instance for the specified file type
func NewReader(fileType FileType) (Reader, error) {
	switch fileType {
	case CSV:
		return NewCSVReader(), nil
	case PDF:
		return NewPDFReader(), nil
	default:
		return nil, ErrUnsupportedFileType
	}
}
