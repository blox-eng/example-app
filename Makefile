.PHONY: build
build:
	-templ generate
	-./tailwindcss -i ./static/css/input.css -o ./static/css/style.css
	-go build -o ./tmp/blox ./cmd/app/main.go
	
.PHONY: dev
dev:
	go build -o ./tmp/blox ./cmd/app/main.go && air

.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	staticcheck ./...

.PHONY: test
test:
	go test -race -v -timeout 30s ./...

