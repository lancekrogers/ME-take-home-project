.PHONY: all test clean build sqlc postgres createdb dropdb migrateup freshdb mock

all: test build

postgres:
	docker run --name me_challenge_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it me_challenge_db createdb --username=root --owner=root me_challenge 

dropdb:
	docker exec -it me_challenge_db dropdb me_challenge 

run: build
	./challenge

build:
	go build -o ./challenge ./cmd

clean:
	rm -f ./challenge

test:
	go test ./...

sqlc:
	sqlc generate

migrateup:
	migrate -path pkg/db/migrations -database

make freshdb: dropdb createdb

mock:
	mockgen -package mockdb -destination pkg/db/mock/repo.go challenge/pkg/db Repo
