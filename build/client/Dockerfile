FROM golang:1.16-buster as builder

ARG service=client

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
# Expecting to copy go.mod and if present go.sum.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY cmd/${service}/. ./
COPY pb ./pb

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ${service} .

FROM alpine:3.12

RUN apk --no-cache add ca-certificates && \
    rm -rf /var/cache/apk/*

COPY --from=builder /app/${service} /app/${service}

ENTRYPOINT  ["/app/${service}"]
