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

## Updating Documentation

If you make changes to the API endpoints, update the annotations in the Go files and regenerate the documentation:

```bash
swag init -g main.go
``` 