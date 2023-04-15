fmt:
	go fmt ./...

deps:
	go get -u ./...
	go mod tidy

build:
	go build -o bin/bidder cmd/bidder/main.go
