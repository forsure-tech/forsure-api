.PHONY: migrate database dev network

dev: network, database

migrate:
ifeq (${new},)
else
	@migrate create -ext sql -dir ./assets/migrations ${new}
endif

ifeq (${run},)
else
	@migrate -source file://assets/migrations -database mysql://root:@tcp\(localhost:3306\)/forsure ${run}
endif

database:
	@docker-compose -f build/dev/docker-compose.yml up
