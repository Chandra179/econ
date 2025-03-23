package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"stock/filereader" // Updated import path based on the module name in go.mod
)

func main() {
	// Create test files
	testDir := "test_files"
	if err := os.MkdirAll(testDir, 0755); err != nil {
		log.Fatalf("Failed to create test directory: %v", err)
	}

	// Create a test CSV file
	csvContent := "id,name,email\n1,John Doe,john@example.com\n2,Jane Smith,jane@example.com"
	csvFilePath := filepath.Join(testDir, "test.csv")
	if err := os.WriteFile(csvFilePath, []byte(csvContent), 0644); err != nil {
		log.Fatalf("Failed to create test CSV file: %v", err)
	}

	// Example 1: Reading a CSV file
	fmt.Println("Example 1: Reading a CSV file")
	csvReader, err := filereader.NewReader(filereader.CSV)
	if err != nil {
		log.Fatalf("Failed to create CSV reader: %v", err)
	}

	csvData, err := csvReader.Read(csvFilePath)
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Print CSV data
	fmt.Println("CSV Data:")
	for _, row := range csvData {
		fmt.Println(row)
	}
	fmt.Println()

	// Example 2: Reading CSV from string
	fmt.Println("Example 2: Reading CSV from string")
	csvContent2 := "product,price,quantity\napple,0.99,10\nbanana,0.59,5"
	csvData2, err := csvReader.ReadFromBytes([]byte(csvContent2))
	if err != nil {
		log.Fatalf("Failed to read CSV from string: %v", err)
	}

	// Print CSV data
	fmt.Println("CSV Data from string:")
	for _, row := range csvData2 {
		fmt.Println(row)
	}
	fmt.Println()

	// Example 3: Reading a PDF file
	fmt.Println("Example 3: Reading a PDF file")

	// Download a sample PDF file if it doesn't exist
	pdfFilePath := filepath.Join(testDir, "sample.pdf")
	if _, err := os.Stat(pdfFilePath); os.IsNotExist(err) {
		fmt.Println("Downloading sample PDF...")
		// Create a simple PDF file using pdfcpu
		err := createSamplePDF(pdfFilePath)
		if err != nil {
			log.Fatalf("Failed to create sample PDF: %v", err)
		}
	}

	pdfReader, err := filereader.NewReader(filereader.PDF)
	if err != nil {
		log.Fatalf("Failed to create PDF reader: %v", err)
	}

	// Cast to PDFReader to access specific settings
	pdfReaderImpl, ok := pdfReader.(*filereader.PDFReader)
	if ok {
		// Configure PDF reader to extract both text and metadata
		pdfReaderImpl.ExtractTextOnly = true
		pdfReaderImpl.ExtractMetadata = true
	}

	pdfData, err := pdfReader.Read(pdfFilePath)
	if err != nil {
		log.Fatalf("Failed to read PDF file: %v", err)
	}

	// Print PDF data
	fmt.Println("PDF Data:")
	for i, row := range pdfData {
		fmt.Println(row)
		// Limit output to avoid flooding the console
		if i > 10 {
			fmt.Println("... more lines ...")
			break
		}
	}
	fmt.Println()

	// Example 4: Using the interface to handle multiple file types
	fmt.Println("Example 4: Using the interface to handle multiple file types")

	// Function that works with any Reader implementation
	processFile := func(filePath string, fileType filereader.FileType) {
		reader, err := filereader.NewReader(fileType)
		if err != nil {
			log.Fatalf("Failed to create reader for %s: %v", fileType, err)
		}

		data, err := reader.Read(filePath)
		if err != nil {
			log.Fatalf("Failed to read file: %v", err)
		}

		fmt.Printf("Data from %s file:\n", fileType)
		for i, row := range data {
			if i < 3 { // Print only first 3 rows to keep output short
				fmt.Println(row)
			} else {
				fmt.Println("...")
				break
			}
		}
		fmt.Println()
	}

	// Process both files with the same function
	processFile(csvFilePath, filereader.CSV)
	processFile(pdfFilePath, filereader.PDF)

	fmt.Println("Cleanup: removing test files")
	os.RemoveAll(testDir)
}

// createSamplePDF creates a simple PDF file using raw data
func createSamplePDF(filepath string) error {
	// Simple PDF structure with some text content
	pdfContent := `%PDF-1.4
1 0 obj
<< /Type /Catalog /Pages 2 0 R >>
endobj
2 0 obj
<< /Type /Pages /Kids [3 0 R] /Count 1 >>
endobj
3 0 obj
<< /Type /Page /Parent 2 0 R /Resources 4 0 R /MediaBox [0 0 612 792] /Contents 5 0 R >>
endobj
4 0 obj
<< /Font << /F1 6 0 R >> >>
endobj
5 0 obj
<< /Length 112 >>
stream
BT
/F1 24 Tf
100 700 Td
(Sample PDF Document for Testing) Tj
0 -50 Td
/F1 12 Tf
(This is a simple PDF file created for testing the PDFReader component.) Tj
ET
endstream
endobj
6 0 obj
<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>
endobj
7 0 obj
<< /Title (Sample PDF) /Author (Stock App) /Subject (PDF Reader Test) /Keywords (test,pdf,reader) /Creator (File Reader Example) /Producer (Go PDF Creator) /CreationDate (D:20230323090000) >>
endobj
xref
0 8
0000000000 65535 f
0000000009 00000 n
0000000058 00000 n
0000000117 00000 n
0000000217 00000 n
0000000258 00000 n
0000000421 00000 n
0000000488 00000 n
trailer
<< /Size 8 /Root 1 0 R /Info 7 0 R >>
startxref
688
%%EOF`

	return os.WriteFile(filepath, []byte(pdfContent), 0644)
}
