build:
	go build -o tictactoe cmd/tictactoe/main.go
run:
	go run cmd/tictactoe/main.go
test:
	go test -v ./...