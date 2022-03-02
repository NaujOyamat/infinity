test:
	go test ./...

test-verbose:
	go test -v ./...

test-coverage:
	go test ./... -cover

test-coverage-html:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out