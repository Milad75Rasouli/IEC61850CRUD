build:
	go build  -o ./bin/IEC61850Api ./cmd/main.go

run: build
	./bin/IEC61850Api