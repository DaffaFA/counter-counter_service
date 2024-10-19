include .env

# ~~~ Database Migrations ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

POSTGRESQLDSN := "postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable"

migrate-up: $(MIGRATE) ## Apply all (or N up) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	echo $(POSTGRESQLDSN); \
	migrate  -database $(POSTGRESQLDSN) -path=migrations up ${NN}

.PHONY: migrate-down
migrate-down: $(MIGRATE) ## Apply all (or N down) migrations.
	@ read -p "How many migration you wants to perform (default value: [all]): " N; \
	migrate  -database $(POSTGRESQLDSN) -path=migrations down ${NN}

.PHONY: migrate-drop
migrate-drop: $(MIGRATE) ## Drop everything inside the database.
	migrate  -database $(POSTGRESQLDSN) -path=migrations drop

.PHONY: migrate-create
migrate-create: $(MIGRATE) ## Create a set of up/down migrations with a specified name.
	@ read -p "Please provide name for the migration: " Name; \
	migrate create -ext sql -dir migrations -seq $${Name}
