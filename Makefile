POSTGRES_PASSWORD ?= mysecretpassword
POSTGRES_PORT ?= 15432

POSTGRES_CONTAINER_NAME := develop-with-skaffold-db

prepare:
	(docker volume ls -q | grep -q $(POSTGRES_CONTAINER_NAME)-vol) || docker volume create $(POSTGRES_CONTAINER_NAME)-vol; \
	docker run -d --rm --name $(POSTGRES_CONTAINER_NAME) --env POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p $(POSTGRES_PORT):5432 -v $(POSTGRES_CONTAINER_NAME)-vol:/var/lib/postgresql/data postgres:12.3-alpine

stop:
	docker stop $(POSTGRES_CONTAINER_NAME)

cleanup: stop
	docker volume rm $(POSTGRES_CONTAINER_NAME)-vol
