FROM golang:1.16-buster as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY cmd/server/. ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server .

FROM alpine:3.12

RUN apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /app/server /app/server

ENTRYPOINT  ["/app/server"]
