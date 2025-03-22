# Stock Market API Documentation

This directory contains the Swagger documentation for the Stock Market API.

## Overview

The Stock Market API provides endpoints for retrieving time series data for various stock symbols.

## Files

- `swagger.yaml` - OpenAPI 3.0 specification in YAML format
- `swagger.json` - OpenAPI 3.0 specification in JSON format
- `docs.go` - Generated Go code for Swagger documentation

## Accessing the Documentation

When the API is running, you can access the Swagger UI at:

```
http://localhost:8080/swagger/index.html
```

This provides an interactive documentation where you can:
- View all available API endpoints
- Try out API calls directly from the browser
- See request/response examples
- View detailed parameter information

## API Endpoints

### Time Series Endpoints
1. **GET /api/v1/timeseries/{symbol}**
   - Gets time series data for a specific stock symbol
   - Supports daily, weekly, and monthly data

2. **GET /api/v1/timeseries/{symbol}/{interval}**
   - Gets intraday time series data with a specific interval
   - Supports intervals: 1min, 5min, 15min, 30min, 60min

## Updating Documentation

If you make changes to the API endpoints, update the annotations in the Go files and regenerate the documentation:

```bash
swag init -g main.go
``` 