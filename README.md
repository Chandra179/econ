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

## Future Extensions

1. Optional features:
   - Historical options: delta, gamma, theta, vega, rho
   - Realtime options
   - Advanced analytics (fixed & sliding) window: correlation, median, variance
   - ETF data
   - Technical indicators: SMA, EMA, WMA, DEMA, TEMA, etc.
