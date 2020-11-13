DOCKER_IMAGE_TAG := 0.0.1

.PHONY: build generate-proto

build:
	go build -o build/pns cmd/main.go

build-docker:
	docker build -t pns .

build-rest:
	go build -o build/rest cmd/rest.go

build-grpc:
	go build -o build/grpc cmd/grpc.go

publish-dockerhub:
	docker build -t bilalekremharmansa/pns:${DOCKER_IMAGE_TAG} .
	docker push bilalekremharmansa/pns:${DOCKER_IMAGE_TAG}

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/notification.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pushservice/push.proto

mock:
	mockgen --source=internal/push/service.go --destination test/push/service_mock.go