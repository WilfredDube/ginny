run:
	go run cmd/main.go

test:
	go test ./... -v

## migrations name=$1: create a new database migration
migration:
	@echo 'Creating migration files for ${name}...'
	migrate create -seq -ext=.sql -dir=./migrations ${name}

up:
	@echo 'Running up migrations...'
	migrate -path=./migrations -database ${SNIPPET_DSN} -verbose up

down:
	@echo 'Running down migration'
	migrate -path=./migrations -database ${SNIPPET_DSN} -verbose down

sqlc:
	sqlc generate