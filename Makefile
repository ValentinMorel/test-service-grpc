.PHONY : build run fresh test clean

test:
	go test ./...

build-deps:
	go mod download

build:
	go build 

run-server:
	go run main.go

run-list: 
	@go run cmd/client/main.go -list

run-whois:
	@go run cmd/client/main.go -whois=$(USER)

clean:
	go clean
