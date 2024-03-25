# Run swag fmt for Swagger files
swag fmt ./...

# Run go fmt for Go files
go fmt ./...

# Run Markdown linting (assuming you have markdownlint-cli installed)
markdownlint -f --config .markdownlint.yaml *.md

# Run Dockerfile linting (assuming you have hadolint installed)
hadolint Dockerfile
