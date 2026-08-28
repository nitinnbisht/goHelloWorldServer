# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-sample-app

# Runtime stage
FROM alpine:3.19
LABEL maintainer="Harness Demo <demo@harness.io>"
WORKDIR /bin
COPY --from=builder /app/go-sample-app .
EXPOSE 8080
ENTRYPOINT ["/bin/go-sample-app"]
