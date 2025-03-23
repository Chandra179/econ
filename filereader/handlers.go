package filereader

import (
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// FileReadResponse defines the response format for file reading operations
// @Description File read response data structure
type FileReadResponse struct {
	Version    string     `json:"version"`
	Timestamp  string     `json:"timestamp"`
	Filename   string     `json:"filename"`
	FileType   string     `json:"fileType"`
	Data       [][]string `json:"data"`
	Error      string     `json:"error,omitempty"`
	RowsCount  int        `json:"rowsCount"`
	ColumnsMax int        `json:"columnsMax"`
}

// ReadFileParams holds parameters for reading files
type ReadFileParams struct {
	FileType FileType `form:"fileType" binding:"required"`
}

// ReadFileFromUpload handles file upload requests and returns parsed data
// @Summary Read and parse an uploaded file
// @Description Reads and parses an uploaded file (CSV or PDF) and returns structured data
// @Tags filereader
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File to be parsed"
// @Param fileType query string true "File type (csv or pdf)" Enums(csv, pdf)
// @Success 200 {object} FileReadResponse "Successful operation"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/filereader/upload [post]
func ReadFileFromUpload(c *gin.Context) {
	// Get the file type parameter
	var params ReadFileParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File type parameter is required (csv or pdf)",
		})
		return
	}

	// Get file from request
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No file provided or invalid file",
		})
		return
	}
	defer file.Close()

	// Create the appropriate reader based on file type
	reader, err := NewReader(params.FileType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Read file data
	data, err := reader.ReadFromReader(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Calculate rows and max columns
	rowsCount := len(data)
	columnsMax := 0
	for _, row := range data {
		if len(row) > columnsMax {
			columnsMax = len(row)
		}
	}

	// Create response
	response := FileReadResponse{
		Version:    "1.0",
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Filename:   header.Filename,
		FileType:   string(params.FileType),
		Data:       data,
		RowsCount:  rowsCount,
		ColumnsMax: columnsMax,
	}

	c.JSON(http.StatusOK, response)
}

// ReadFileFromBytes handles requests with file data in the request body
// @Summary Read and parse file data from request body
// @Description Reads and parses file data (CSV or PDF) from the request body and returns structured data
// @Tags filereader
// @Accept octet-stream
// @Produce json
// @Param fileType query string true "File type (csv or pdf)" Enums(csv, pdf)
// @Param filename query string false "Original filename (for reference only)"
// @Param body body []byte true "Raw file content"
// @Success 200 {object} FileReadResponse "Successful operation"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/filereader/readbytes [post]
func ReadFileFromBytes(c *gin.Context) {
	// Get the file type parameter
	var params ReadFileParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File type parameter is required (csv or pdf)",
		})
		return
	}

	// Get filename (optional)
	filename := c.Query("filename")
	if filename == "" {
		filename = "unknown"
	}

	// Read request body
	fileData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request body",
		})
		return
	}

	// Create the appropriate reader based on file type
	reader, err := NewReader(params.FileType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Read file data
	data, err := reader.ReadFromBytes(fileData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Calculate rows and max columns
	rowsCount := len(data)
	columnsMax := 0
	for _, row := range data {
		if len(row) > columnsMax {
			columnsMax = len(row)
		}
	}

	// Create response
	response := FileReadResponse{
		Version:    "1.0",
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
		Filename:   filename,
		FileType:   string(params.FileType),
		Data:       data,
		RowsCount:  rowsCount,
		ColumnsMax: columnsMax,
	}

	c.JSON(http.StatusOK, response)
}
