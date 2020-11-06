
.PHONY: test
test:
	cd test/push; go test .

mock:
	mockgen --source=internal/push/service.go --destination test/push/service_mock.go