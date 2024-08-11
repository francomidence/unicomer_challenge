# Holiday Service API

This is a simple Holiday Service API built with Golang and Gin framework. It fetches holiday data from an external API and provides filtering options based on holiday type and date range.

## Features
- Fetch holiday data from an external API.
- Filter holidays by type (e.g., Civil, Religioso).
- Filter holidays by date range.
- Rate limiting to prevent abuse.
- CORS support for cross-origin requests.
- Automatic request ID generation for easier logging and tracing.

## Middlewares
1. **CORS Middleware**: Allows cross-origin requests and controls what methods and headers are allowed.
2. **Rate Limiter Middleware**: Limits the number of requests a client can make within a specified timeframe.
3. **Request ID Middleware**: Generates a unique request ID for each incoming request, aiding in tracing and logging.

## Logging
Logging includes a request id for each request along with useful information on the results.

## Docker
`docker build -t unicomer_challenge .`
`docker run -p 8080:8080 unicomer_challenge`
