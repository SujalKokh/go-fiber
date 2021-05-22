include .env
MIGRATE=docker-compose exec web migrate -path=migration -database "mysql://${DBUsername}:${DBPassword}@tcp(${DBHost}:${DBPort})/${DBName}" -verbose

dev:
		gin --appPort 8090 -i run main.go
migrate-up:
		$(MIGRATE) up
migrate-down:
		$(MIGRATE) down 
force:
		@read -p  "Which version do you want to force?" VERSION; \
		$(MIGRATE) force $$VERSION

goto:
		@read -p  "Which version do you want to migrate?" VERSION; \
		$(MIGRATE) goto $$VERSION

drop:
		$(MIGRATE) drop

create:
		@read -p  "What is the name of migration?" NAME; \
		${MIGRATE} create -ext sql -dir migration  $$NAME

.PHONY: migrate-up migrate-down force goto drop create