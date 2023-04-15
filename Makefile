fmt:
	go fmt ./...

build:
	go build -o bin/bidder cmd/bidder/main.go

run:
	bin/bidder
