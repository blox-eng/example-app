.PHONY: run_db
run_db:
	@docker run --hostname=957e5b1bfd30 --env=POSTGRES_PASSWORD=admin --env=PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/lib/postgresql/13/bin --env=GOSU_VERSION=1.17 --env=LANG=en_US.utf8 --env=PG_MAJOR=13 --env=PG_VERSION=13.14-1.pgdg120+2 --env=PGDATA=/var/lib/postgresql/data --volume=blox_eng:/var/lib/postgresql/data --volume=/var/lib/postgresql/data -p 5432:5432 --restart=no --runtime=runc -d postgres:13

.PHONY: create_tests
create_tests:
	find . -type f -name '*_test.go' -exec rm {} +
	gotests -all -w ./*
