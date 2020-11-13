DOCKER_IMAGE_TAG := 0.0.1

.PHONY: build generate-proto

define build_go_file
	GOOS=windows go build -o build/$(1).exe cmd/$(2).go
	GOOS=linux go build -o build/$(1) cmd/$(2).go
endef

build:clean-build
	$(call build_go_file,pns,main)

build-docker:
	docker build -t pns .

build-rest:clean-build
	$(call build_go_file,rest,rest)

build-grpc: clean-build
	$(call build_go_file,grpc,grpc)

clean-build:
	rm -rf build/

publish-dockerhub:
	docker build -t bilalekremharmansa/pns:${DOCKER_IMAGE_TAG} .
	docker push bilalekremharmansa/pns:${DOCKER_IMAGE_TAG}

generate-proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/notification.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/pushservice/push.proto

mock:
	mockgen --source=internal/push/service.go --destination test/push/service_mock.go