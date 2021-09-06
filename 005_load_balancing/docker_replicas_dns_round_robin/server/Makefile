setup:
	docker build -t app-server .

start:
	docker run --name app-server-demo --publish 9000:9000 --env ID=app1 --detach --rm app-server

test:
	curl http://localhost:9000

stop:
	docker stop app-server-demo

teardown:
	docker rmi `docker images --quiet app-server`