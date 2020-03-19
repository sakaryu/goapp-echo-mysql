.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up

.PHONY: lint
lint:
	docker-compose run --rm app golint -set_exit_status ./...

.PHONY: init-db
init-db: migrate-db fixtures-db

.PHONY: migrate-db
migrate-db:
	docker-compose run --rm -v $$PWD:/mnt/ mysql /mnt/bin/migrate.sh

.PHONY: fixtures-db
fixtures-db:
	docker-compose run --rm -v $$PWD:/mnt/ mysql /mnt/bin/fixtures.sh
