BINARY_NAME=kosyncsrv
 
build:
	CGO_ENABLED=1 go build -o ${BINARY_NAME} *.go

run:
	CGO_ENABLED=1 go build -o ${BINARY_NAME} *.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

.PHONY: test
test:
	go test ./... -race -cover .

.PHONY: lint
lint:
	golangci-lint run