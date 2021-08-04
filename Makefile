build:
	docker-compose -f docker-compose.yml build
up:
	docker-compose -f docker-compose.yml up
up-d:
	docker-compose -f docker-compose.yml up -d
down:
	docker-compose -f docker-compose.yml down
logs:
	docker-compose -f docker-compose.yml logs --tail=100
