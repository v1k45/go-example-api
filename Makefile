build:
	go build -o bin/$(shell basename $(PWD)) .

server: build
	./bin/$(shell basename $(PWD)) server

migrations: build
	./bin/$(shell basename $(PWD)) migrations up

dev:
	air
