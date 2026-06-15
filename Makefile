include .env

migrateup:
	migrate -path db/migrations -database "$(DATABASE_URL)" up

migratedown:
	migrate -path db/migrations -database "$(DATABASE_URL)" down

migrationforce:
	migrate -path db/migrations -database "$(DATABASE_URL)" force 1

sqlc:
	sqlc generate