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

