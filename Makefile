.PHONY: pb clean

all: clean pb server client

server:
#CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server cmd/server/main.go
	go build -ldflags="-s -w" -o server cmd/server/*.go

client:
#CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o client cmd/client/main.go
	go build -ldflags="-s -w" -o client cmd/client/*.go

pb:
	protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    pb/run.proto

clean:
	rm -f client server
