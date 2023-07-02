test:
	go test ./... -count=1 -p 4

test_coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

.PHONY: test test_coverage
