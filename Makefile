.PHONY: version
version:
	go version
	protoc --version

.PHONY: generate
generate:
	protoc --go_out=pkg --go-grpc_out=pkg api/reviews.proto

.PHONY: run
run:
	go run ./cmd/reviews/main.go

.PHONY: test
test:
	cd test && go test ./... -v -tags test -count=1
