all: format_code download_deps test_with_coverage build update_docs

format_code:
	@echo "Formatting code"
	@go fmt ./...

download_deps:
	@echo "Downloading dependencies"
	@go mod download
	@go mod verify

update_docs:
	@echo "Generating documentation"
	@cd docs && go run doc_generator.go

build:
	# Build docker image (the image will have the compiled binaries in it, refer the Dockerfile)
	docker build -t cli_builder .

	# Create a container from that image
	CONTAINER_NAME=cli_container_topson
	# Remove if it already exists
	docker rm $CONTAINER_NAME || true
	docker create --name $CONTAINER_NAME cli_builder

	# Extract the built binaries from the image to local disk
	docker cp $CONTAINER_NAME:/go/cli/dist dist

test:
	@echo "Running tests.."
	@go test -v ./...
	@echo "Tests done !"

test_with_coverage:
	@echo "Running tests.."
	@go test -v ./... -coverprofile=cover.out -coverpkg=./...
	@echo "Tests Passed, writing coverage report to cover.html"
	@go tool cover -html=cover.out -o cover.html
	@echo "Tests done !"

.PHONY: all build_bin test_with_coverage update_docs format_code test build download_deps
