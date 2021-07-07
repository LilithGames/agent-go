PROTOC_DIR := pkg/transfer

.PHONY: autogen
autogen:
	@cd ${PROTOC_DIR} && protoc --go_out=. --go-grpc_out=. *.proto

.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o ./examples/bin/agent.test ./examples/hello

.PHONY: build-image
build-image: build
	@docker-compose -f examples/deploy/docker-compose.yaml build hello