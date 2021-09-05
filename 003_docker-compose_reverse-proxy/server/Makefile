setup:
	docker build -t docker-compose-server .

start:
	docker run --name docker-compose-server-demo --publish 9000:9000 --detach --rm docker-compose-server

test:
	curl http://localhost:9000

stop:
	docker stop docker-compose-server-demo

teardown:
	docker rmi `docker images --quiet docker-compose-server`