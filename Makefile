.PHONY: run produce consume db clean migrate

run: db produce consume

stop: clean

produce:
	@go run cmd/producer/main.go --length=7 --amount=500 --path storage/tokens

consume:
	@go run cmd/consumer/main.go --path storage/tokens

db:
	@docker run -d \
	--name prodcon-postgres \
	-e POSTGRES_HOST_AUTH_METHOD=trust \
	-e POSTGRES_USER=user \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=prodcon \
	-p 5433:5432 \
	postgres:13.3-alpine

migrate:
	docker run -v migrations:/migrations \
	--network host \
	migrate/migrate \
	-path=/migrations \
	-database 'postgres://user@localhost:5433/prodcon?sslmode=disable' \
	up 2

clean:
	@docker stop prodcon-postgres
	@docker rm prodcon-postgres
	@rm storage/tokens

#psql -d prodcon -U user -h localhost -p 5432 -W
