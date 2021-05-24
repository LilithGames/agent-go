PROTOC_DIR := pkg/transfer

.PHONY: autogen
autogen:
	@cd ${PROTOC_DIR} && protoc --go_out=. --go-grpc_out=. *.proto