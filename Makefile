build:
	@go build -o bin/go-users-api cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./bin/go-users-api
