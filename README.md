# Stock Market Data API

This API provides access to stock market time series data with versioning support.

## Quick Start

1. Clone the repository
2. Copy `.env.example` to `.env` and update with your Alpha Vantage API key
3. Run the server: `go run main.go`
4. Access the Swagger UI at `http://localhost:8080/swagger/index.html`

## API Documentation

The API is documented using OpenAPI 3.0 Specification. You can view the interactive documentation in several ways:

1. **Using the built-in Swagger UI**:
   ```bash
   go run main.go
   # Visit http://localhost:8080/swagger/index.html in your browser
   ```

2. **Using Docker Compose with Swagger UI**:
   ```bash
   docker-compose up swagger-ui
   # Visit http://localhost:8081 in your browser
   ```

3. **Using Online Swagger Editor**:
   - Go to [Swagger Editor](https://editor.swagger.io/)
   - Import the `docs/swagger.yaml` file

For more details about the API documentation, see the [docs/README.md](docs/README.md) file.

## API Endpoints

### Time Series Data

The API uses versioning to ensure compatibility as it evolves.

#### Current version: v1

**Base URL**: `/api/v1/timeseries`

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/api/v1/timeseries/:symbol` | GET | Get time series data for a specific symbol |
| `/api/v1/timeseries/:symbol/:interval` | GET | Get time series data with a specific interval (1min, 5min, 15min, 30min, 60min) |

**Query Parameters**:

- `function`: Time series function (default: TIME_SERIES_DAILY)
  - Options: TIME_SERIES_INTRADAY, TIME_SERIES_DAILY, TIME_SERIES_WEEKLY, TIME_SERIES_MONTHLY
- `outputsize`: Output size (default: compact)
  - Options: compact, full
- `datatype`: Data type (default: json)
  - Options: json, csv
- `adjusted`: Whether to return adjusted data (default: false)
- `extended_hours`: Whether to include extended hours data (for intraday only, default: false)
- `month`: Monthly data in YYYY-MM format (for intraday only)
- `apikey`: Your Alpha Vantage API key (optional if set in .env)

**Example Request**:

```
GET /api/v1/timeseries/AAPL?function=TIME_SERIES_DAILY&outputsize=compact
```

**Example Response**:

```json
{
  "version": "1.0",
  "timestamp": "2024-03-22T10:30:00Z",
  "symbol": "AAPL",
  "data": {
    "Meta Data": {
      "1. Information": "Daily Prices (open, high, low, close) and Volumes",
      "2. Symbol": "AAPL",
      "3. Last Refreshed": "2024-03-21",
      "4. Output Size": "Compact",
      "5. Time Zone": "US/Eastern"
    },
    "Time Series (Daily)": {
      "2024-03-21": {
        "1. open": "177.5000",
        "2. high": "178.5200",
        "3. low": "176.8300",
        "4. close": "177.2800",
        "5. volume": "58769727"
      },
      // Additional data points...
    }
  }
}
```

## Future Extensions

1. Optional features:
   - Historical options: delta, gamma, theta, vega, rho
   - Realtime options
   - Advanced analytics (fixed & sliding) window: correlation, median, variance
   - ETF data
   - Technical indicators: SMA, EMA, WMA, DEMA, TEMA, etc.