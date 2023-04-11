build:
	@go build -o bin/jwt_auth

run: build
	@./bin/jwt_auth

test:
	@go test -v ./...
