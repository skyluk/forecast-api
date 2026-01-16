# Weather API

A lightweight Go-based REST API that provides weather forecast information by querying the National Weather Service API.

## Overview

This application acts as a wrapper around the [National Weather Service API](https://www.weather.gov/documentation/services-web-api), providing a simple interface to retrieve forecast data for any geographic coordinate. The API accepts latitude and longitude as parameters and returns detailed weather forecasts.

## Prerequisites

- **Go**: Version 1.25.1 or later
- **Internet connection**: Required to fetch data from the National Weather Service API

## Installation

1. Clone the repository:
```bash
git clone https://github.com/skyluk/weather-api.git
cd weather-api
```

2. Download dependencies:
```bash
go mod download
```

## Running the Application

Start the server with:

```bash
go run ./cmd/main.go
```

The API will start on `http://localhost:8080`

You should see:
```
Starting weather forecast API on port 8080
```

## API Usage

### Get Forecast

**Endpoint:** `GET /api/v1/forecast/:coordinate`

**Parameters:**
- `coordinate`: Latitude and longitude in the format `latitude,longitude`

**Example Request:**
```bash
curl -X GET http://localhost:8080/api/v1/forecast/40.18443,-105.1467
```

This example retrieves the forecast for Boulder, Colorado (latitude: 40.18443, longitude: -105.1467).

**Response:**
Returns JSON-formatted weather forecast data from the National Weather Service API.

## Project Structure

- `cmd/main.go` - Application entry point and route definitions
- `internal/adapters/weather/` - Weather API adapter for communicating with NWS API
- `internal/server/` - HTTP server and request handlers
- `go.mod` / `go.sum` - Go module dependencies

## Dependencies

- [httprouter](https://github.com/julienschmidt/httprouter) - High-performance HTTP router

## Notes

- The National Weather Service API base URL is currently hardcoded. In a production system, this should be configurable via environment variables or a configuration file.
- The server uses standard HTTP (not HTTPS) on port 8080. For production deployment, use HTTPS via `http.ListenAndServeTLS()` or proxy through a web server.
- The API relies on external data from the National Weather Service; ensure your application can reach `https://api.weather.gov`.

## License

[Add your license information here]
