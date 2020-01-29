test:
	go test ./...

unit_test:
	go test `go list ./... | grep -v e2e_test`