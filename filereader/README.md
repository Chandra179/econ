# File Reader Package

A Go package providing a flexible interface for reading and parsing different file formats. Currently supports CSV and PDF files using Go standard packages.

## Features

- Common interface for reading different file formats
- Support for CSV files using the Go standard `encoding/csv` package
- Basic support for PDF text extraction using Go standard packages
- Multiple input methods: file path, byte slice, or io.Reader
- Flexible error handling

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    "stock/filereader"
)

func main() {
    // Create a CSV reader
    csvReader, err := filereader.NewReader(filereader.CSV)
    if err != nil {
        log.Fatalf("Failed to create CSV reader: %v", err)
    }
    
    // Read CSV file
    data, err := csvReader.Read("path/to/file.csv")
    if err != nil {
        log.Fatalf("Failed to read CSV file: %v", err)
    }
    
    // Process data
    for _, row := range data {
        fmt.Println(row)
    }
    
    // Create a PDF reader
    pdfReader, err := filereader.NewReader(filereader.PDF)
    if err != nil {
        log.Fatalf("Failed to create PDF reader: %v", err)
    }
    
    // Read PDF file
    pdfData, err := pdfReader.Read("path/to/file.pdf")
    if err != nil {
        log.Fatalf("Failed to read PDF file: %v", err)
    }
    
    // Process PDF data
    for _, row := range pdfData {
        fmt.Println(row)
    }
}
```

### Reading from a Byte Slice

```go
// Read CSV from a byte slice
csvContent := "id,name\n1,John\n2,Jane"
data, err := csvReader.ReadFromBytes([]byte(csvContent))
if err != nil {
    log.Fatalf("Failed to read CSV from bytes: %v", err)
}
```

### Reading from an io.Reader

```go
import (
    "os"
    "stock/filereader"
)

// Open a file
file, err := os.Open("path/to/file.csv")
if err != nil {
    log.Fatalf("Failed to open file: %v", err)
}
defer file.Close()

// Create a CSV reader
csvReader, err := filereader.NewReader(filereader.CSV)
if err != nil {
    log.Fatalf("Failed to create CSV reader: %v", err)
}

// Read CSV from io.Reader
data, err := csvReader.ReadFromReader(file)
if err != nil {
    log.Fatalf("Failed to read CSV from reader: %v", err)
}
```

### CSV Reader Configuration

The CSV reader provides configuration options for customizing the parsing behavior:

```go
// Get a CSV reader with custom settings
csvReader := filereader.NewCSVReader()
csvReader.Comma = ';'           // Use semicolon as separator
csvReader.Comment = '#'         // Lines starting with # are comments
csvReader.TrimLeadingSpace = true // Trim leading whitespace in fields

// Use the configured reader
data, err := csvReader.Read("path/to/file.csv")
```

## Notes on PDF Support

The PDF reader implementation in this package uses Go standard packages and provides basic text extraction capabilities. It is primarily designed for simple text-based PDFs. For more advanced PDF processing, consider integrating a specialized PDF library.

## Extending with New File Types

To add support for a new file type:

1. Define a new constant in `interface.go`
2. Create a new reader struct that implements the `Reader` interface
3. Update the `NewReader` function to return your new reader for the new file type

The interface-based design makes it easy to add new file formats while maintaining compatibility with existing code. 