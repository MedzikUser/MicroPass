BINARY_NAME=micropass

all: release

build:
	go build -o $(BINARY_NAME) main.go

release:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME)-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o $(BINARY_NAME)-linux-arm64 main.go
	GOOS=freebsd GOARCH=amd64 go build -o $(BINARY_NAME)-freebsd-amd64 main.go
	GOOS=freebsd GOARCH=arm64 go build -o $(BINARY_NAME)-freebsd-arm64 main.go

run:
	env MICROPASS_DEBUG=true go run main.go

clean:
	go clean
	rm -f $(BINARY_NAME)*
