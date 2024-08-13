build:
	go build -o ./bin/dsa ./cmd/dsa

run:
	./bin/dsa

all: build run
