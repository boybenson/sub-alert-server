run: air

build:
	go build -o bin/app main.go

dev:
	air

test:
	go test ./...

clean:
	rm -rf bin

swagger:	
	swag init -g main.go -o docs
