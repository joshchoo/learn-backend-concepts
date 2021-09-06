setup:
	docker build -t docker-compose-server .

start:
	docker run --name=docker-compose-server-demo --publish=9000:9000 --env-file=../local.env --rm docker-compose-server

logs:
	docker logs docker-compose-server-demo

test:
	curl http://localhost:9000

stop:
	docker stop docker-compose-server-demo

teardown:
	docker rmi `docker images --quiet docker-compose-server`