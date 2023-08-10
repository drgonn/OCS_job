

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
