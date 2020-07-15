POSTGRES_PASSWORD ?= mysecretpassword
POSTGRES_PORT ?= 15432

POSTGRES_CONTAINER_NAME := develop-with-skaffold-db

prepare:
	docker run -d --rm --name $(POSTGRES_CONTAINER_NAME) --env POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p $(POSTGRES_PORT):5432 -v $(CURDIR)/db:/var/lib/postgresql/date postgres:12.3-alpine

cleanup:
	docker stop $(POSTGRES_CONTAINER_NAME)
