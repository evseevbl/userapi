# generate stub for `storage`
.PHONY: mock
mock:
	minimock -i github.com/evseevbl/userapi/internal/app/userapi.storage -o ./internal/app/userapi


.PHONY: tests
tests:
	go test ./...
	go test ./... -tags=integration