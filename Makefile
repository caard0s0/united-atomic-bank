postgres:
	docker run --name postgres15.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15.2 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres15.2 dropdb bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up	

migrateup1: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up	1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test: 
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go atomic-bank/db/sqlc Store

_PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc server mock
