.PHONY: build
build:
	@templ generate
	@go build -o blox ./cmd/api/main.go
	@npx tailwindcss -i ./cmd/web/styles/main.css -o ./cmd/web/styles/tailwind.css

.PHONY: install
install:
	"$(CURDIR)/scripts/install.sh"
.PHONY: develop
develop:
	@air .
	@templ generate --watch

.PHONY: run
run:
	@go run cmd/api/main.go

.PHONY: run_db
run_db:
	"$(CURDIR)/scripts/run_database.sh"

.PHONY: create_tests
create_tests:
	"$(CURDIR)/scripts/create_tests.sh"

.PHONY: docs
docs:
	"$(CURDIR)/scripts/docs.sh"

.PHONY: style
style:
	"$(CURDIR)/scripts/style.sh"

.PHONY: test
test:
	"$(CURDIR)/scripts/test.sh"

