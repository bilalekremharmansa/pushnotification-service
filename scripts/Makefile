.PHONY: build generate-proto

build:
	go build -o build/rest cmd/rest.go

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/notification.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pushservice/push.proto

mock:
	mockgen --source=internal/push/service.go --destination test/push/service_mock.go