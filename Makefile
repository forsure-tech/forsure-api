.PHONY: migrate database

migrate:
ifeq (${new},)
else
	@migrate create -ext sql -dir ./assets/migrations ${new}
endif

database:
	@docker-compose -f build/dev/docker-compose.yml up
