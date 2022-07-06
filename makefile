run:
	go run ./cmd/main.go
build-arm64:
	env GOOS=linux GOARCH=arm64 go build -o ./build/taskmanager ./cmd/*.go
build-local:
	go build -o ./build/taskmanager.local ./cmd/*.go
