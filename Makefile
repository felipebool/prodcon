.PHONY: run produce consume db clean migrate

stop: clean

produce:
	@go run cmd/producer/main.go --length=7 --amount=1000 --path storage/tokens

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

clean:
	@docker stop prodcon-postgres
	@docker rm prodcon-postgres
	@rm storage/tokens

#psql -d prodcon -U user -h localhost -p 5432 -W
