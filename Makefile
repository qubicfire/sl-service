COMPOSE = docker compose --env-file .env.production

.PHONY: up down restart build logs db migrate migrate-down migrate-status create-migration

up:
	$(COMPOSE) up -d

up-db:
	$(COMPOSE) up -d postgres

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) restart

logs:
	$(COMPOSE) logs -f app

logs-db:
	$(COMPOSE) logs -f postgres

build:
	$(COMPOSE) up -d --build

db:
	$(COMPOSE) exec postgres psql -U sl_user -d sl_service

migrate:
	$(COMPOSE) --profile migrate run --rm migrate up

migrate-down:
	$(COMPOSE) --profile migrate run --rm migrate down

migrate-status:
	$(COMPOSE) --profile migrate run --rm migrate status

create-migration:
	goose -dir migrations create $(name) sql

