DB_URL=postgres://kevin:kube@localhost:5432/dbname?sslmode=disable

install-tools:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.26.0
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0

migrate-up:
	migrate -path internal/db/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path internal/db/migrations -database "$(DB_URL)" down

migrate-create:
	@if [ -z "$(MIGRATION_NAME)" ]; then \
		echo "Error: MIGRATION_NAME is required"; \
		exit 1; \
	fi
	migrate create -ext sql -dir internal/db/migrations -seq $(MIGRATION_NAME)

sqlc:
	sqlc generate -f internal/db/sqlc.yaml

run:
	go run cmd/server/main.go


