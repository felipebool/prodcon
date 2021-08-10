.PHONY: produce consume db clean test

stop: clean

produce:
	@go run cmd/producer/main.go

consume:
	@go run cmd/consumer/main.go

db:
	@docker run -d \
	--name prodcon-postgres \
	-e POSTGRES_HOST_AUTH_METHOD=trust \
	-e POSTGRES_USER=user \
	-e POSTGRES_PASSWORD=password \
	-e POSTGRES_DB=prodcon \
	-p 5433:5432 \
	postgres:13.3-alpine \
	-c shared_buffers=256MB \
	-c max_connections=200 \
	-c work_mem=32MB \

clean:
	@docker stop prodcon-postgres
	@docker rm prodcon-postgres
	@rm storage/tokens

test:
	@go test ./...

# ------------------------------------------------------------------------------
# To check the saved data inside the database,
# copy and paste the following commands:
# log into container: docker exec -it prodcon-postgres bash
# connect to database: psql -d prodcon -U user -h localhost -p 5432
# ------------------------------------------------------------------------------
