postgres:
	docker run --name postgres15.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine

createdb:
	docker exec -it postgres15.2 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres15.2 dropdb bank

migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose up	

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bank?sslmode=disable" -verbose down

_PHONY: postgres createdb dropdb migrateup migratedown
