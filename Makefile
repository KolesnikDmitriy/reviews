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
	cd test && go test ./... -tags test -count=1

.PHONY: local
local: export BASE_URL=localhost:50051
local: test