.PHONY: run
run:
	go run cmd/creator-applications/main.go

.PHONY: build
build:
	go build cmd/creator-applications/main.go

.PHONY: gen
gen:
	protoc -I=api --go_out=internal/pb --go-grpc_out=internal/pb api/* --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative