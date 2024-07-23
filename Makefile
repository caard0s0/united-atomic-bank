postgres:
	docker run --name postgres16.3 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.3-alpine

createdb:
	docker exec -it postgres16.3 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres16.3 dropdb bank

migrateup: 
	migrate -path internal/database/migrations -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up	

migrateup1: 
	migrate -path internal/database/migrations -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up	1

migratedown:
	migrate -path internal/database/migrations -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path internal/database/migrations -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test: 
	go test -v -cover -short ./...

server:
	go run cmd/main.go

mock:
	mockgen -package mockdb -destination internal/database/mock/store.go github.com/caard0s0/vanguard-server/internal/database/sqlc Store

_PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server mock
