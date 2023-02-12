DB_URL=postgresql://root:4kb1cXvaFYNh0ejSFd8a@simple-bank.ck18pryaeny8.ap-southeast-1.rds.amazonaws.com:5432/simple_bank

network:
	docker network create bank-network

#################
###  DATABASE ###
#################
postgres:
	docker run --name postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank


##################
###  MIGRATION ###
##################
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1


#####################
###  SQLC GenCode ###
#####################
sqlc:
	sqlc generate


#####################
###  Test go ########
#####################
test:
	go test -v -cover ./...


#####################
###  Server go ######
#####################
server:
	go run main.go


mock:
	mockgen -package mockdb -destination db/mock/store.go backend-master-class-simple-bank/db/sqlc Store

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 db_docs db_schema sqlc test server mock proto evans redis
