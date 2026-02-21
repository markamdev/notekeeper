default:
	@echo "This is default target. Use one of targets listed below:"
	@echo "proto - build prtobuf files"
	@echo "client - build client binary"
	@echo "server - build server binary"

.PHONY: proto client server all clean

proto:
	@echo "Compiling protobuf files"
	protoc --go_out=. --go_opt=paths=source_relative \
    	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
    	proto/notekeeper/notekeeper.proto

client: proto
	@echo "Compiling client binary"
	go build -o client ./cmd/client/main.go

server: proto
	@echo "Compiling server binary"
	go build -o server ./cmd/server/...

all: proto client server

clean:
	@echo "Cleaning repo"
	rm -f client server

install-deps:
	@echo "Installing dependencies:"
	@echo "- protoc-gen-go"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@echo "- protoc-gen-go-grpc"
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest