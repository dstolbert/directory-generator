unit:
	go test -v -coverprofile=coverage.out `go list ./... | egrep -v '(/test|/mocks|/client|/cmd|/gateway)'`
	go tool cover -func coverage.out

mocks:
	components .

# generic run command
run:
	go run ./cmd/generator/

# run in dev mode
dev:
	go run ./cmd/dev/