package filereader

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

// CSVReader implements the Reader interface for CSV files
type CSVReader struct {
	Comma            rune // Field delimiter (comma by default)
	Comment          rune // Comment character (disabled by default)
	FieldsPerRecord  int  // Number of fields per record (-1 means unspecified)
	LazyQuotes       bool // Allow lazy quotes
	TrimLeadingSpace bool // Trim leading space in fields
}

// NewCSVReader creates a new CSVReader with default settings
func NewCSVReader() *CSVReader {
	return &CSVReader{
		Comma:            ',',
		Comment:          0,
		FieldsPerRecord:  -1,
		LazyQuotes:       false,
		TrimLeadingSpace: false,
	}
}

// Read parses the CSV file at the given path
func (r *CSVReader) Read(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, WrapError(err, "failed to open CSV file")
	}
	defer file.Close()

	return r.ReadFromReader(file)
}

// ReadFromBytes parses CSV data from a byte slice
func (r *CSVReader) ReadFromBytes(data []byte) ([][]string, error) {
	return r.ReadFromReader(bytes.NewReader(data))
}

// ReadFromReader parses CSV data from an io.Reader
func (r *CSVReader) ReadFromReader(reader io.Reader) ([][]string, error) {
	csvReader := csv.NewReader(reader)

	// Apply configuration
	csvReader.Comma = r.Comma
	csvReader.Comment = r.Comment
	csvReader.FieldsPerRecord = r.FieldsPerRecord
	csvReader.LazyQuotes = r.LazyQuotes
	csvReader.TrimLeadingSpace = r.TrimLeadingSpace

	// Read all records
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, WrapError(err, "failed to parse CSV data")
	}

	return records, nil
}
