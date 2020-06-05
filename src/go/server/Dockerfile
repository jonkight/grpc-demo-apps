# Build and test the app
FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go test -v ./tests/... \
    && go build -o main . \
    && mkdir /dist \
    && cp main /dist

# Create the deploy artifact
FROM scratch

COPY --from=builder /dist/main /app

ENTRYPOINT ["/app/main"]
