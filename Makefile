BINARY_NAME=go-hello-http

test:
	go test ./...

build:
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-amd64 main.go
	GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}-darwin-arm64 main.go
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go

clean:
	go clean
	rm -f ${BINARY_NAME}-darwin-amd64
	rm -f ${BINARY_NAME}-darwin-arm64
	rm -f ${BINARY_NAME}-linux
