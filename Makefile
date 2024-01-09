BINARY_NAME = campsite_p2p

build:
	@go build -o bin/$(BINARY_NAME) -v ./p2p.go

clean:
	@rm -f bin/$(BINARY_NAME)

run: build
	@./bin/$(BINARY_NAME)

test:
	@go test -v ./...

# Path: campsite_p2p.go
