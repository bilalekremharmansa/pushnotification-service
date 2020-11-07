.PHONY: build

build:
	go build -o build/rest cmd/rest.go

mock:
	mockgen --source=internal/push/service.go --destination test/push/service_mock.go