FROM golang:1.21-alpine AS builder

WORKDIR /src

# Download modules
COPY go.mod .
RUN go mod download

# Copy source and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /garage ./

FROM alpine:3.18
RUN apk add --no-cache ca-certificates

COPY --from=builder /garage /garage

EXPOSE 8090

ENTRYPOINT ["/garage"]
