build:
	go build  -o ./bin/server ./cmd/server/main.go

run: build
	./bin/server