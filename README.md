# HelloWorldGoServer

A simple Go HTTP server demo for Harness CI/CD integration. Features graceful shutdown, query parameter handling, and comprehensive test coverage.

## Prerequisites

- Go 1.21 or higher
- Docker (optional, for containerization)

## Building

### Local Build
```bash
go build -o go-sample-app
```

### Docker Build
```bash
docker build -t go-sample-app:latest .
```

The multi-stage Dockerfile optimizes layer caching for faster builds:
- **Build stage**: Uses `golang:1.21-alpine` to compile
- **Runtime stage**: Minimal `alpine:3.19` image (~15MB)

## Running the App

### Locally
```bash
./go-sample-app
# Output: Starting Server
```

### With Docker
```bash
docker run -p 8080:8080 go-sample-app:latest
```

### Testing Requests
```bash
curl http://localhost:8080?name=Alice
# Output: Hello, Alice

curl http://localhost:8080
# Output: Hello, Guest
```

## Running Tests

### Quick Test Run
```bash
go test -v
```

### Generate JUnit XML Report
```bash
go install github.com/jstemmer/go-junit-report/v2@latest
go test -v 2>&1 | go-junit-report -set-exit-code > test-results.xml
```

### Generate Coverage Reports

#### HTML Coverage Report
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

#### LCOV Format (for CI integrations)
```bash
go install github.com/jandelgado/gcov2lcov@latest
go test -coverprofile=coverage.out
gcov2lcov -infile=coverage.out -outfile=coverage.lcov
```

### Combined: Tests + JUnit + Coverage
```bash
go test -v -coverprofile=coverage.out 2>&1 | tee test-output.log
go-junit-report -set-exit-code -in test-output.log -out test-results.xml
gcov2lcov -infile=coverage.out -outfile=coverage.lcov
```

## Test Files

- **hello_server_test.go**: Original tests for greeting functionality
- **utils_test.go**: String utility and format tests
- **handlers_test.go**: HTTP handler tests with test recorder
- **validation_test.go**: Input validation and UTF-8 compliance tests
- **edge_cases_test.go**: Tests for special characters, long names, whitespace
- **integration_test.go**: Multi-request and query parsing tests

## Harness CI Integration

Example `.harness/` pipeline stages:

```yaml
# Build
- step:
    type: Run
    spec:
      command: go build -o go-sample-app

# Test with JUnit
- step:
    type: Run
    spec:
      command: |
        go install github.com/jstemmer/go-junit-report/v2@latest
        go test -v 2>&1 | go-junit-report -set-exit-code > test-results.xml
      reports:
        type: JUnit
        spec:
          paths:
            - test-results.xml

# Coverage with LCOV
- step:
    type: Run
    spec:
      command: |
        go install github.com/jandelgado/gcov2lcov@latest
        go test -coverprofile=coverage.out
        gcov2lcov -infile=coverage.out -outfile=coverage.lcov
      reports:
        type: Codecov
        spec:
          paths:
            - coverage.lcov

# Docker Build
- step:
    type: BuildAndPush
    spec:
      dockerfile: Dockerfile
      repo: <your-registry>/go-sample-app
      tags:
        - latest
        - <+pipeline.executionId>
```

## Architecture

- **Server**: Gorilla Mux router, graceful shutdown on SIGINT/SIGTERM
- **Handler**: Query parameter parsing with default value ("Guest")
- **Logging**: Standard library logging
- **Timeout**: 10s read/write timeouts

## License

MIT
