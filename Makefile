ENV_FILE := ./configs/.env
ifneq (,$(wildcard $(ENV_FILE)))
	include $(ENV_FILE)
	export $(shell sed 's/=.*//' $(ENV_FILE))
endif

run:
	docker compose -f deployments/docker-compose.yml --env-file ./configs/.env up --build -d

migration_up:
	migrate -path migrations// -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5433/$(POSTGRES_DB)?sslmode=disable" -verbose up

migration_down:
	migrate -path migrations// -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5433/$(POSTGRES_DB)?sslmode=disable" -verbose down

gen_doc: 
	swag init -g main.go -d cmd/app,internal/app,internal/database,internal/models,internal/services,internal/transport

.PHONY: run migration_up migration_down gen_doc
