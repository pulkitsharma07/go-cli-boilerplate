all: format_code test build_bin update_docs


format_code:
	@echo "Formatting code"
	@go fmt ./...

update_docs:
	@echo "Generating documentation"
	@cd docs && go run doc_generator.go

build_bin:
	@echo "Building basecommand"
	@go build -ldflags "-s -w" -o build/basecommand
	@echo "Build done"

test:
	@echo "Running tests.."
	@go test -v ./... -coverprofile=cover.out -coverpkg=./...
	@echo "Tests Passed, writing coverage report to cover.html"
	@go tool cover -html=cover.out -o cover.html
	@echo "Tests done !"

.PHONY: all build_bin test update_docs format_code
