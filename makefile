run:
	go run ./cmd/cli/main.go

build-cli:
	go build -o ./build/blockchain.exe ./cmd/cli/...

test:
	go test ./...