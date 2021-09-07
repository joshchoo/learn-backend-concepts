start:
	docker-compose up --build

inspect:
	docker-compose ps

test:
	curl http://localhost:8080/

scale-up:
	docker-compose up --scale server=5 --detach

scale-down:
	docker-compose up --scale server=1 --detach

stop:
	docker-compose down