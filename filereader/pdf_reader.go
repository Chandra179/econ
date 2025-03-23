package filereader

import (
	"bytes"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// PDFReader implements the Reader interface for PDF files
type PDFReader struct {
	// Configuration options
	ExtractMetadata bool
	ExtractTextOnly bool
}

// NewPDFReader creates a new PDFReader with default settings
func NewPDFReader() *PDFReader {
	return &PDFReader{
		ExtractMetadata: true,
		ExtractTextOnly: true,
	}
}

// Read parses the PDF file at the given path
func (r *PDFReader) Read(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, WrapError(err, "failed to open PDF file")
	}
	defer file.Close()

	return r.ReadFromReader(file)
}

// ReadFromBytes parses PDF data from a byte slice
func (r *PDFReader) ReadFromBytes(data []byte) ([][]string, error) {
	return r.ReadFromReader(bytes.NewReader(data))
}

// ReadFromReader parses PDF data from an io.Reader
func (r *PDFReader) ReadFromReader(reader io.Reader) ([][]string, error) {
	// Create a temporary file to store the PDF data
	tmpFile, err := os.CreateTemp("", "pdf-*.pdf")
	if err != nil {
		return nil, WrapError(err, "failed to create temporary file")
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Copy content from reader to temp file
	if _, err := io.Copy(tmpFile, reader); err != nil {
		return nil, WrapError(err, "failed to write PDF data to temporary file")
	}

	// Ensure data is written to disk
	if err := tmpFile.Sync(); err != nil {
		return nil, WrapError(err, "failed to sync temporary file")
	}

	// Close and reopen to ensure file is properly flushed
	tmpFile.Close()

	var result [][]string

	// Extract text content from the PDF using file-based API
	if r.ExtractTextOnly {
		// Extract content to a temporary directory
		tmpDir, err := os.MkdirTemp("", "pdf-extract-*")
		if err != nil {
			return nil, WrapError(err, "failed to create temporary directory")
		}
		defer os.RemoveAll(tmpDir)

		// Use the file-based API to extract content
		conf := model.NewDefaultConfiguration()
		err = api.ExtractContentFile(tmpFile.Name(), tmpDir, nil, conf)
		if err != nil {
			// Try plain text extraction as fallback
			if err := api.ExtractPagesFile(tmpFile.Name(), tmpDir, nil, conf); err == nil {
				// Read the extracted pages
				contentFiles, err := os.ReadDir(tmpDir)
				if err == nil {
					extractTextFromFiles(contentFiles, tmpDir, &result)
				}
			}

			// If we still don't have results, try one more approach
			if len(result) == 0 {
				result = append(result, []string{"Text extraction fallback"})
				textBytes, err := extractRawText(tmpFile.Name())
				if err == nil && len(textBytes) > 0 {
					processExtractedText(textBytes, &result)
				}
			}
		} else {
			// Read the extracted content
			contentFiles, err := os.ReadDir(tmpDir)
			if err != nil {
				return nil, WrapError(err, "failed to read content directory")
			}

			// Process extracted content files
			extractTextFromFiles(contentFiles, tmpDir, &result)
		}
	}

	// Extract metadata if requested
	if r.ExtractMetadata {
		metadata, err := extractMetadata(tmpFile.Name())
		if err != nil {
			// Don't fail the entire operation if metadata extraction fails
			result = append(result, []string{"Metadata extraction failed: " + err.Error()})
		} else {
			// Add a separator before metadata
			if len(result) > 0 {
				result = append(result, []string{"----- Metadata -----"})
			}

			for k, v := range metadata {
				result = append(result, []string{"Metadata", k, v})
			}
		}
	}

	if len(result) == 0 {
		return [][]string{{"No content extracted from PDF"}}, nil
	}

	return result, nil
}

// extractTextFromFiles processes extracted text files and adds their content to the result
func extractTextFromFiles(files []os.DirEntry, baseDir string, result *[][]string) {
	for _, file := range files {
		if !file.IsDir() && (strings.HasSuffix(file.Name(), ".txt") || strings.HasSuffix(file.Name(), ".content")) {
			data, err := os.ReadFile(baseDir + "/" + file.Name())
			if err == nil {
				processExtractedText(data, result)
			}
		}
	}
}

// processExtractedText cleans and formats extracted PDF text
func processExtractedText(data []byte, result *[][]string) {
	text := string(data)

	// Clean up PDF markup
	text = cleanPDFMarkup(text)

	// Split text into lines and add them to the result
	lines := strings.Split(text, "\n")
	var currentParagraph strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			// End of paragraph
			if currentParagraph.Len() > 0 {
				*result = append(*result, []string{currentParagraph.String()})
				currentParagraph.Reset()
			}
		} else {
			// Continue building paragraph
			if currentParagraph.Len() > 0 {
				currentParagraph.WriteString(" ")
			}
			currentParagraph.WriteString(line)
		}
	}

	// Add the last paragraph if it's not empty
	if currentParagraph.Len() > 0 {
		*result = append(*result, []string{currentParagraph.String()})
	}
}

// cleanPDFMarkup removes PDF-specific markup and formatting codes
func cleanPDFMarkup(text string) string {
	// Remove PDF operators and commands
	operatorPattern := regexp.MustCompile(`\b(BT|ET|Tj|TJ|Td|TD|Tm|T\*|Tc|Tw|Tz|TL|Tf|Tr|Ts)\b`)
	text = operatorPattern.ReplaceAllString(text, "")

	// Remove parentheses around text
	text = strings.ReplaceAll(text, "( ", "")
	text = strings.ReplaceAll(text, " )", "")
	text = strings.ReplaceAll(text, "(", "")
	text = strings.ReplaceAll(text, ")", "")

	// Remove square brackets
	text = strings.ReplaceAll(text, "[", "")
	text = strings.ReplaceAll(text, "]", "")

	// Remove font declarations
	fontPattern := regexp.MustCompile(`/F\d+\s+\d+(\.\d+)?\s+Tf`)
	text = fontPattern.ReplaceAllString(text, "")

	// Remove positioning commands
	posPattern := regexp.MustCompile(`\d+(\.\d+)?\s+\d+(\.\d+)?\s+Td`)
	text = posPattern.ReplaceAllString(text, "")

	// Clean up whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")
	text = strings.TrimSpace(text)

	return text
}

// extractRawText tries to extract text from PDF using raw extraction
func extractRawText(filename string) ([]byte, error) {
	// Create a temporary directory for extraction
	tmpDir, err := os.MkdirTemp("", "pdf-text-extract-*")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	// Extract text using pdfcpu
	conf := model.NewDefaultConfiguration()
	if err := api.ExtractPagesFile(filename, tmpDir, nil, conf); err != nil {
		return nil, err
	}

	// Combine all extracted text
	var allText bytes.Buffer
	files, err := os.ReadDir(tmpDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			data, err := os.ReadFile(tmpDir + "/" + file.Name())
			if err == nil {
				allText.Write(data)
				allText.WriteString("\n")
			}
		}
	}

	return allText.Bytes(), nil
}

// extractMetadata extracts important metadata from a PDF file
func extractMetadata(filename string) (map[string]string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Use pdfcpu API to extract metadata
	conf := model.NewDefaultConfiguration()
	properties, err := api.Properties(file, conf)
	if err != nil {
		return nil, err
	}

	// Return all properties as they're already filtered
	return properties, nil
}
