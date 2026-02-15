.PHONY: run build test templ dev clean

run:
	go run cmd/server/main.go

build:
	go build -o bin/goblog cmd/server/main.go

test:
	go test ./... -v

templ:
	templ generate

dev:
	templ generate && go run cmd/server/main.go

clean:
	rm -rf bin/ tmp/
