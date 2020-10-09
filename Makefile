.PHONY: version
version:
	go version
	protoc --version

.PHONY: generate
generate:
	protoc --go_out=pkg --go-grpc_out=pkg api/reviews.proto
