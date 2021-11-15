## help: Show current help
help: Makefile
	@echo "Choose a command run:"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'

## go-format: Format and simplify all source files
go-format:
	@printf "Format Go source files... " && gofmt -s -w ./go && echo "Done!"

## go-lint: Check source code by linters
go-lint:
	@printf "Checking via golangci-lint... " && GO111MODULE=off golangci-lint run ./go/... && echo "Done!"

## go-test: Run Golang tests
go-test:
	@echo "Running Golang tests... " && GO111MODULE=off go test ./go/... && echo "Done!"

## rust-test: Run Rust tests
rust-test:
	@echo "Running Rust tests... " && cargo test --manifest-path=./rust/Cargo.toml --tests && echo "Done!"

.PHONY: help go-format go-lint go-test rust-test
