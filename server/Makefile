.PHONY: build
build: build-protos build-go-code

.PHONY: build-protos
build-protos:
	protoc \
		--proto_path=../protos \
		--go_out=. --go_opt=module=github.com/DaikoneKisu/recycle-it/server \
		--go-grpc_out=. --go-grpc_opt=module=github.com/DaikoneKisu/recycle-it/server \
		../protos/**/*

.PHONY: build-go-code
build-go-code:
	go mod tidy
	go fmt
	go build