DB_URL=postgresql://root:secret@localhost:5432/flash_sale?sslmode=disable

# Start database (enter the deploy directory to start compose)
up:
	docker-compose up -d

# Close database
down:
	docker-compose down -v

# Create new migration file
new_migration:
	migrate create  -ext sql -dir Server/db/migration -seq init_schema

# Up migration
migrateup:
	migrate -path Server/db/migration -database "$(DB_URL)" -verbose up 
# Down migrate
migratedown:
	migrate -path Server/db/migration -database "$(DB_URL)" -verbose down

# Execute the next migration file
migrateup1:
	migrate -path Server/db/migration -database "$(DB_URL)" -verbose up 1

# Revoke the most recent migration
migratedown1:
	migrate -path Server/db/migration -database "$(DB_URL)" -verbose down 1

# generate sql
sqlc:
	sqlc generate




.PHONY: up dowv new_migration migrateup migratedown migrateup1 migratedown1 sqlc

