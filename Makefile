

up:
	docker-compose up -d

down:
	docker-compose down

stop:
	docker-compose stop

shell:
	docker exec -it ocs bash

install:
	go mod tidy 

run:
	go run main.go


redis:
	docker exec -it redis_ocs redis-cli

sql:
	docker exec -it postgres_ocs psql 

psql:
	docker exec -it postgres_ocs psql -a "user=root password=root dbname=ocs"