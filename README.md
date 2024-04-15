# Average Number Service

## Description

Lightweight express app which calls a configured endpoint periodically and calculates a cumulative average of random numbers.

### Behaviours and assumptions

- The configured handler assumes a response type as detailed here https://csrng.net/documentation/csrng-lite/.

- If any call to the endpoint returns an error, the application exits
  - This is to ensure any application runner/container can restart if required
- The endpoint is subject to rate limiting, but in this use case it is not predicted to occur, as such any error of this nature is subject to the above.
  - If requirements favoured uptime instead of calling every second, then a retry mechanism with backoff would be introduced to adhere to the rate limit.
- The latest number retrieved from the service is stored in memory, but instead of storing _all_ collected numbers in memory the application calculates a cumulative average to avoid memory bloat over time.

## Run

To start the service, run

```
go run main.go
```

## Testing

There are unit tests to cover the functionality and an integration test to verify calling the express server.

To run

```
go test
```
