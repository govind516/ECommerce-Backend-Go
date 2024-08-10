build: 
	@go build -o bin/E-Commerce cmd/main.go

test: 
	@go test -v ./...

run: build
	@./bin/E-Commerce